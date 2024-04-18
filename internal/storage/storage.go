package storage

import (
	"calculator_final/internal/entities"
	"fmt"
	"reflect"
)

// хранилище ЗАЛОГИНЕННЫХ пользователей не выполняет особо важных функций
type Storage struct {
	UsersLogged []entities.User
}

func New() *Storage {
	users := make([]entities.User, 0)
	return &Storage{users}
}

func (s *Storage) AddUser(user entities.User) error {
	for _, u := range s.UsersLogged {
		// проверяем нет ли уже этого пользователя
		if reflect.DeepEqual(u, user) {
			return fmt.Errorf("user %v already exists", user.Username)
		}
	}
	s.UsersLogged = append(s.UsersLogged, user)
	return nil
}

func (s *Storage) GetUser(Username string) (*entities.User, error) {
	for _, user := range s.UsersLogged {
		if user.Username == Username {
			return &user, nil
		}
	}
	return &entities.User{}, fmt.Errorf("user %v not found", Username)
}

// func (s *Storage) EditUserStatus(ID, status string) error {
// 	for _, user := range s.UsersLogged {
// 		user := user
// 		if user.Id == ID {
// 			user.Logged = status == "logged"
// 			logger.Info(fmt.Sprintf("user %v status changed to 'logged'", user.Username))
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("user with ID %v not found", ID)
// }
