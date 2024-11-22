package main

import (
	"composer_vue/backend"
	"composer_vue/backend/app/compose"
	"composer_vue/cmd"
	"composer_vue/database/seeder"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db := cmd.InitDB()
	seeder.Seed(db)

	repo := compose.NewGormComposeRepository(db)
	composeService := compose.NewGormCompose(repo)

	app := backend.NewApp(composeService)

	err := wails.Run(&options.App{
		Title:  "composer_vue",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
