package main

import (
	"embed"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// compute covers dir so we can mount it via the asset server middleware
	cfgDir, _ := os.UserConfigDir()
	appDir := filepath.Join(cfgDir, "PenguinTunes")
	coversDir := filepath.Join(appDir, "covers")

	// middleware to serve /covers/* from the filesystem
	coversMiddleware := func(next http.Handler) http.Handler {
		fs := http.StripPrefix("/covers/", http.FileServer(http.Dir(coversDir)))
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.HasPrefix(r.URL.Path, "/covers/") {
				fs.ServeHTTP(w, r)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "PenguinTunes",
		Width:  1400,
		MinWidth:  1024,
		Height: 768,
		MinHeight: 420,
		AssetServer: &assetserver.Options{
			Assets: assets,
			Middleware: assetserver.ChainMiddleware(coversMiddleware),
		},
		Frameless: true,
		BackgroundColour: &options.RGBA{R: 23, G: 23, B: 25, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
