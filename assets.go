package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func getCoversMiddleware() func(http.Handler) http.Handler {
	cfgDir, _ := os.UserConfigDir()
	appDir := filepath.Join(cfgDir, "PenguinTunes")
	coversDir := filepath.Join(appDir, "covers")

	return func(next http.Handler) http.Handler {
		fs := http.StripPrefix("/covers/", http.FileServer(http.Dir(coversDir)))
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/covers/") {
				fs.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// getAssetsMiddleware serves audio files directly from the filesystem using an absolute path.
// Usage examples:
//   `/asset?path=/home/user/music/song.mp3`
//   `/asset/%2Fhome%2Fuser%2Fmusic%2Fsong.mp3`  (encoded absolute path in URL)
func getAssetsMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/asset" && !strings.HasPrefix(r.URL.Path, "/asset") {
				next.ServeHTTP(w, r)
				return
			}
			// Extract path from query param or from URL path
			var p string
			if q := r.URL.Query().Get("path"); q != "" {
				p = q
			} else if strings.HasPrefix(r.URL.Path, "/asset/") {
				enc := strings.TrimPrefix(r.URL.Path, "/asset/")
				if dec, err := url.PathUnescape(enc); err == nil {
					p = dec
				} else {
					http.Error(w, "invalid path", http.StatusBadRequest)
					return
				}
			} else {
				http.Error(w, "missing path", http.StatusBadRequest)
				return
			}

			// Ensure absolute path
			if !filepath.IsAbs(p) {
				http.Error(w, "path must be absolute", http.StatusBadRequest)
				return
			}

			// Check file exists and is a regular file
			fi, err := os.Stat(p)
			if err != nil || fi.IsDir() {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}

			// CORS: allow requests from the renderer (including wails:// origin)
			origin := r.Header.Get("Origin")
			allowOrigin := "*"
			if origin != "" {
				allowOrigin = origin
			}

			// Handle preflight requests
			if r.Method == http.MethodOptions {
				w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Range, Icy-Metadata, Content-Type")
				w.Header().Set("Access-Control-Max-Age", "86400")
				w.WriteHeader(http.StatusNoContent)
				return
			}

			// Set CORS headers for actual requests and expose range-related headers
			w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Expose-Headers", "Accept-Ranges, Content-Range, Content-Length")

			fmt.Printf("Serving asset file: %s (origin=%s)\n", p, origin)

			// Serve the file (supports range requests)
			http.ServeFile(w, r, p)
		})
	}
}
