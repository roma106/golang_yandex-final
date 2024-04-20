package html

import (
	"calculator_final/internal/databases"
	"calculator_final/internal/logger"
	"calculator_final/internal/storage"
	"calculator_final/internal/utils"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func LoadHtmlPage(w http.ResponseWriter, r *http.Request, name string, data any) {
	// загрузка страниц
	tmpl := template.Must(template.ParseFiles(fmt.Sprintf("frontend/%s", name)))
	tmpl.Execute(w, data)
}

func LoadImg(w http.ResponseWriter, r *http.Request) {
	// загружка картинок
	filename := strings.Split(r.URL.Path, "imgs/")[1]
	http.ServeFile(w, r, fmt.Sprintf("frontend/imgs/%s", filename))
}

func LoadStaticFile(w http.ResponseWriter, r *http.Request) {
	// загрузка всякой дичи по типу style.css и jsники
	filename := strings.Split(r.URL.Path, "static/")[1]
	http.ServeFile(w, r, fmt.Sprintf("frontend/static/%s", filename))
}

func LoadCalculatorPage(w http.ResponseWriter, r *http.Request, storage *storage.Storage, db *sql.DB) {
	// получаем имя пользователя из строки запроса
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "no username provided", http.StatusForbidden)
		return
	}
	// если user залогинен (есть в массиве залогиненных челиков)- выдаем ему страничку
	user, err := storage.GetUser(username)
	if err == nil && user != nil {
		// получили ID пользователя - достаем токен из кеша и возвращаем в header
		token, err := storage.GetToken(user.ID)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to get token from storage. Error: %v", err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Authorization", token.Token)

		// подгружаем выражения пользователя из БД его таблицы
		exprs, err := databases.GetExpressions(db, user.ID)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to get expressions from database. Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		LoadHtmlPage(w, r, "calc.html", exprs)
		return
	}
	// если user не найден - выдаем ему страничку регистрации
	logger.Info("Someone wasn't registered and so was sent to register page. Username: " + username)
	err = utils.RespondWithJSON(w, http.StatusForbidden, "go to register")
	if err != nil {
		logger.Error("error while responding with JSON:" + err.Error())
	}
}
