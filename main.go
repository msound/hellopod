package main

import (
	"log"
	"log/slog"
	"net/http"

	"github.com/msound/hellopod/pkg/app"
	"github.com/msound/hellopod/pkg/config"
)

const APP_VERSION = "v1.0.1"

func main() {
	slog.Info("Starting hellopod", "version", APP_VERSION)
	appConfig := config.LoadConfig()
	app := app.NewApp(appConfig)

	log.Fatal(http.ListenAndServe(":"+appConfig.Port, app.GetHandler()))
}
