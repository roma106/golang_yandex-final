package databases

import (
	"calculator_final/internal/entities"
	"calculator_final/internal/logger"
	"database/sql"
	"fmt"
)

func AddUser(db *sql.DB, user entities.User) error {
	_, err := db.Exec(`INSERT INTO "calculator-users" ("Username", "Password", "LastLoggedIn") VALUES ($1, $2, $3)`, user.Username, user.Password, user.LastLogged)
	if err != nil {
		logger.Error("Failed to add user to database: " + err.Error())
		return err
	}
	logger.Info(fmt.Sprintf("User %s added to database", user.Username))
	return nil
}

func GetUsers(db *sql.DB) ([]entities.User, error) {
	rows, err := db.Query(`SELECT * FROM "calculator-users"`)
	if err != nil {
		logger.Error("Failed to select users from database: " + err.Error())
		return nil, err
	}
	defer rows.Close()
	users := []entities.User{}
	for rows.Next() {
		user := entities.User{}
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.LastLogged)
		if err != nil {
			logger.Error("Failed to scan users from database: " + err.Error())
			continue
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByUsername(db *sql.DB, Username string) (entities.User, error) {
	// Функция для нахождения userа по имени
	users, err := GetUsers(db)
	lastLoggedUser := entities.User{}
	if err != nil {
		return entities.User{}, err
	}
	for _, user := range users {
		if user.Username == Username {
			if lastLoggedUser.Username == "" || user.LastLogged.After(lastLoggedUser.LastLogged) {
				// проверка на случай если несколько пользователей с одинаковым именем(берем того кто регался последним)
				lastLoggedUser = user
			}
		}
	}
	if lastLoggedUser.Username != "" {
		return lastLoggedUser, nil
	}
	return entities.User{}, fmt.Errorf("user with Username %s not found", Username)
}

func EditUserByID(db *sql.DB, ID string, user entities.User) error {
	_, err := db.Exec(`
	UPDATE "calculator-users" SET "Username"=$1, "Password"=$2, "LastLoggedIn"=$3 WHERE "ID"=$4
	`, user.Username, user.Password, user.LastLogged, ID)
	if err != nil {
		return err
	}
	logger.Info(fmt.Sprintf("User %s edited in database", user.Username))
	return nil
}

func CheckUsername(db *sql.DB, username string) (bool, error) {
	// Функция для проверки существования такого имени пользователя (для корректной регистрации)
	users, err := GetUsers(db)
	if err != nil {
		return false, err
	}
	for _, user := range users {
		if user.Username == username {
			return true, nil
		}
	}
	return false, nil
}
