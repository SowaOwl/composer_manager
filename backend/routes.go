package backend

import (
	"composer_vue/backend/app/compose"
	"composer_vue/backend/app/model"
	"context"
)

type App struct {
	ctx            context.Context
	ComposeService *compose.ComposeService
}

func NewApp(composeService *compose.ComposeService) *App {
	return &App{
		ComposeService: composeService,
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAllTypes() []model.PublicType {
	return a.ComposeService.GetAllTypes()
}

func (a *App) GetAllContainers() []model.Container {
	return a.ComposeService.GetAllContainers()
}
