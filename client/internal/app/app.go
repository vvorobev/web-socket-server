package app

import (
	httpapp "ws-client/internal/app/httpApp"
)

type App struct {
	HTTPServer *httpapp.App
}

func New(port int) *App {
	httpApp := httpapp.New(port)

	return &App{HTTPServer: httpApp}
}
