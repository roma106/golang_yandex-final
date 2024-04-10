package app

import (
	"fmt"
	"net/http"

	"jwt_lesson/internal/auth"
	"jwt_lesson/internal/html"
	"jwt_lesson/internal/logger"
)

// func (a *App) Stop(ctx context.Context) error {
// 	// err := a.server.Shutdown(ctx)
// }

type App struct {
	server *http.Server
}

func New() (*App, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { html.LoadHtmlPage(w, r, "login.html") })
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) { html.LoadHtmlPage(w, r, "register.html") })
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) { html.LoadStaticFile(w, r) })
	mux.HandleFunc("/imgs/", func(w http.ResponseWriter, r *http.Request) { html.LoadImg(w, r) })
	mux.HandleFunc("/auth-login", auth.LogIn)
	mux.HandleFunc("/auth-reg", auth.Reg)

	app := new(App)
	app.server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	return app, nil
}

func (a *App) Run() error {
	logger.Info("Starting server...")
	fmt.Println("Your link for webpage: http://localhost:8080/login")
	err := a.server.ListenAndServe()
	if err != nil {
		return err
	}
	logger.Info("Server stopped")
	return nil
}
