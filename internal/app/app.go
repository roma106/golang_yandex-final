package app

import (
	"database/sql"
	"fmt"
	"net/http"

	"calculator_final/internal/auth"
	"calculator_final/internal/databases"
	"calculator_final/internal/expressions"
	"calculator_final/internal/html"
	"calculator_final/internal/logger"
	"calculator_final/internal/storage"
)

// func (a *App) Stop(ctx context.Context) error {
// 	// err := a.server.Shutdown(ctx)
// }

type App struct {
	server        *http.Server
	usersDB       *sql.DB
	expressionsDB *sql.DB
}

func New() (*App, error) {
	app := new(App)

	storage := storage.New()

	// Подключаем сервер

	mux := http.NewServeMux()

	mux.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		html.LoadCalculatorPage(w, r, storage, app.expressionsDB)
	})
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { html.LoadHtmlPage(w, r, "login.html", nil) })
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) { html.LoadHtmlPage(w, r, "register.html", nil) })
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) { html.LoadStaticFile(w, r) })
	mux.HandleFunc("/imgs/", func(w http.ResponseWriter, r *http.Request) { html.LoadImg(w, r) })
	mux.HandleFunc("/auth-login", func(w http.ResponseWriter, r *http.Request) { auth.LogIn(w, r, storage, app.usersDB) })
	mux.HandleFunc("/auth-reg", func(w http.ResponseWriter, r *http.Request) { auth.Reg(w, r, storage, app.usersDB, app.expressionsDB) })
	mux.HandleFunc("/new-expr", func(w http.ResponseWriter, r *http.Request) {
		expressions.NewExpression(w, r, app.usersDB, app.expressionsDB)
	})

	app.server = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Подключаем Базы данных

	dbusers := databases.ConnectToDB("calculator-users")
	err := databases.InitUsers(dbusers)
	if err == nil {
		app.usersDB = dbusers
	}

	dbexprs := databases.ConnectToDB("calculator-expressions")
	app.expressionsDB = dbexprs

	return app, nil
}

func (a *App) Run() error {
	logger.Info("Starting server...")
	fmt.Println("Your link for webpage: http://localhost:8080/login")
	defer logger.Info("Database connection closed")

	err := a.server.ListenAndServe()
	if err != nil {
		return err
	}
	logger.Info("Server stopped")
	return nil
}
