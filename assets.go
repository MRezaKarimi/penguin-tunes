package main

import (
	"net/http"
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
