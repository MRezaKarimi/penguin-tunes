package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "PenguinTunes",
		Width:  1400,
		MinWidth:  1024,
		Height: 768,
		MinHeight: 420,
		AssetServer: &assetserver.Options{
			Assets: assets,
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
