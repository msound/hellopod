package main

import (
	"log"
	"net/http"

	"github.com/msound/hellopod/pkg/app"
	"github.com/msound/hellopod/pkg/config"
)

func main() {
	appConfig := config.LoadConfig()
	app := app.NewApp(appConfig)

	log.Fatal(http.ListenAndServe(":"+appConfig.Port, app.GetIndexHandler()))
}
