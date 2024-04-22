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
	app := NewApp()

	err := wails.Run(&options.App{
		Title:         "CoreCraft",
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        true,
		CSSDragProperty:  "--wails-draggable",
		CSSDragValue:     "drag",
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		//-------------------ПРИ СБОРКИ В Realse УБРАТЬ КОМЕНТ--------------------
		//Windows: &windows.Options{
		//WebviewUserDataPath: executableDir + "\\data\\",
		//},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
