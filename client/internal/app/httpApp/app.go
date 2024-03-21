package httpapp

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
)

type App struct {
	httpServer http.Handler
	port       int
}

func New(port int) *App {

	mux := http.NewServeMux()
	controller(mux)

	return &App{httpServer: mux, port: port}
}

func (app *App) Run() error {
	op := "httpapp.Run"
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	fmt.Printf("http server started %s\n", l.Addr().String())

	if err := http.Serve(l, app.httpServer); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (app *App) MustRun() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func controller(mux *http.ServeMux) {
	mux.Handle("GET /public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("views/index.html"))
		templ.ExecuteTemplate(w, "index", nil)
	})
}
