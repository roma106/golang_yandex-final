package auth

import (
	"calculator_final/internal/databases"
	"calculator_final/internal/entities"
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

func Reg(w http.ResponseWriter, r *http.Request, storage *storage.Storage, usersdb, expressionsdb *sql.DB) {
	// обрабатываем запрос
	user := entities.User{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// проверяем нет ли пользователя с таким же именем

	if userexists, err := databases.CheckUsername(usersdb, user.Username); userexists || err != nil {
		w.WriteHeader(http.StatusForbidden)
		logger.Error(fmt.Sprintf("user %s tried to register, but such usersname already exists", user.Username))
		return
	}

	// Заносим в БД, если все ок, то прокидываем в storage.AddUser с ID из БД
	user.LastLogged = utils.FormatTime(time.Now())

	err = databases.AddUser(usersdb, user)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to add user to database. Error: %v", err))
		return
	}
	// Получаем ID для user из БД
	userFromDB, err := databases.GetUserByUsername(usersdb, user.Username)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get user by Username from database. Error: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// добавляем пользователю ID
	user.ID = userFromDB.ID

	// на всякий случай проверяем все ли сошлось
	userFromDB.LastLogged = utils.FormatTime(userFromDB.LastLogged)
	if !reflect.DeepEqual(user, userFromDB) {
		logger.Error(fmt.Sprintf("user from registration doesn't match user from database. Registration user: %v, User, added to database: %v", user, userFromDB))
		return
	}

	// создаем таблицу выражений для пользователя
	err = databases.CreateExpressionsTable(expressionsdb, user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// добавляем в список залогиненных пользователей
	err = storage.AddUser(user)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to add user to storage. Error: %v", err))
		return
	}
	logger.Info(fmt.Sprintf("user %v added to a storage", user.Username))
	logger.Info(fmt.Sprintf("user %v successfully registered", user.Username))
}
