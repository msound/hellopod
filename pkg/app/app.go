package app

import (
	"html/template"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/msound/hellopod/pkg/config"
)

type App struct {
	config *config.Config
	t      *template.Template
}

func NewApp(config *config.Config) *App {
	tpl := getTemplate()
	t, err := template.New("web").Parse(tpl)
	if err != nil {
		panic("cannot parse template")
	}
	return &App{config: config, t: t}
}

func (app *App) getIndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		remoteHost, _, _ := net.SplitHostPort(r.RemoteAddr)
		data := map[string]any{
			"Pod":           app.config.Pod,
			"Host":          r.Host,
			"Path":          r.URL.Path,
			"RemoteIP":      remoteHost,
			"XForwardedFor": r.Header.Get("X-Forwarded-For"),
		}
		app.t.Execute(w, data)
	})
}

func (app *App) GetHandler() http.Handler {
	loggingHandler := handlers.CombinedLoggingHandler(os.Stdout, app.getIndexHandler())

	return loggingHandler
}

func getTemplate() string {
	return `
<!DOCTYPE html>
<html>
<head><title>Hello Pod!</title></head>
<body>
<h1>Hello Pod!</h1>
<p>Pod: {{ .Pod }}</p>
<p>Host: {{ .Host }}</p>
<p>Path: {{ .Path }}</p>
<p>RemoteIP: {{ .RemoteIP }}</p>
<p>X-Forwarded-For: {{ .XForwardedFor }}</p>
</body>
</html>`
}
