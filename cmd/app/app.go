package app

import "context"

type App struct {
	
}

func InitApp(_ context.Context) *App {
	app := &App{}
	
	return app
}

func (a *App) RunApp(_ context.Context) {
}
