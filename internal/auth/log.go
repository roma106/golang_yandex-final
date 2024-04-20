package auth

import (
	"calculator_final/internal/databases"
	"calculator_final/internal/entities"
	"calculator_final/internal/jwt"
	"calculator_final/internal/logger"
	"calculator_final/internal/storage"
	"calculator_final/internal/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"
)

func LogIn(w http.ResponseWriter, r *http.Request, storage *storage.Storage, db *sql.DB) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user := entities.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.LastLogged = utils.FormatTime(time.Now())

	// чекаем БД(+проверка пароля), если user есть и все ок, то смотрим его ID и прокидываем в storage.AddUser
	userFromDB, err := databases.GetUserByUsername(db, user.Username)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get user by Username from database. Error: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// проверка пароля
	if user.Password != userFromDB.Password {
		logger.Error(fmt.Sprintf("user %v tryed to log in, but entered wrong password", user.Username))
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// добавляем user'у ID
	user.ID = userFromDB.ID

	// на всякий случай проверяем все ли сошлось
	userFromDB.LastLogged = user.LastLogged // обновляем время последнего входа для проверки
	if !reflect.DeepEqual(user, userFromDB) {
		logger.Error(fmt.Sprintf("user from login doesn't match user from database. Login user: %v; User, from database: %v", user, userFromDB))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// все супер - добавляем user в массив залогиненных
	err = storage.AddUser(user)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to add user to storage. Error: %v", err))
	}
	// обновляем время последнего входа в БД
	err = databases.EditUserByID(db, user.ID, user)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to edit user in database. Error: %v", err))
	}

	// обновим токен
	token, err := jwt.CreateJWT(user.Username)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to create new token. Error: %v. User: %v", err, user.Username))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Добавляем токен в кеш, чтобы получить его после редиректа
	err = storage.AddToken(user.ID, token)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to add token to storage. Error: %v. User: %v", err, user.Username))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, string("/calc?username="+user.Username), http.StatusTemporaryRedirect)
	logger.Info(fmt.Sprintf("user %v successfully logged in", user.Username))
}
