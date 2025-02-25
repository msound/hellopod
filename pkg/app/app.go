package app

import (
	"fmt"
	"net/http"

	"github.com/msound/hellopod/pkg/config"
)

type App struct {
	config config.Config
}

func NewApp(config config.Config) *App {
	return &App{config: config}
}

func (app *App) GetIndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		output := fmt.Sprintf("Hello pod!\nHostname: %s\n", app.config.Hostname)
		w.Write([]byte(output))
	})
}
