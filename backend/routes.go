package backend

import (
	"composer_vue/backend/app/compose"
	"composer_vue/backend/app/model"
	"context"
)

type App struct {
	ctx               context.Context
	ComposeController *compose.ControllerImpl
}

func NewApp(composeController *compose.ControllerImpl) *App {
	return &App{
		ComposeController: composeController,
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetAllTypes() []model.PublicType {
	return a.ComposeController.GetAllTypes()
}

func (a *App) GetAllContainers() []model.Container {
	return a.ComposeController.GetAllContainers()
}

func (a *App) GenerateDockerCompose() {
	err := a.ComposeController.GenerateDockerCompose()
	if err != nil {
		panic(err)
	}
}
