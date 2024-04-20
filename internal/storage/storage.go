package storage

import (
	"calculator_final/internal/entities"
	"fmt"
	"reflect"
)

// хранилище ЗАЛОГИНЕННЫХ пользователей не выполняет особо важных функций, кэш для jwt нужен для передачи токена при редиректе на страницу калькулятора
type Storage struct {
	UsersLogged []entities.User
	JWTokens    []entities.JWT
}

func New() *Storage {
	users := make([]entities.User, 0)
	tokens := make([]entities.JWT, 0)
	return &Storage{users, tokens}
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

func (s *Storage) AddToken(userId string, tokenString string) error {
	if userId == "" || tokenString == "" {
		return fmt.Errorf("user id or token string is empty")
	}
	token := entities.JWT{Token: tokenString, UserID: userId}
	s.JWTokens = append(s.JWTokens, token)
	return nil
}

func (s *Storage) GetToken(userId string) (*entities.JWT, error) {
	for _, token := range s.JWTokens {
		if token.UserID == userId {
			return &token, nil
		}
	}
	return &entities.JWT{}, fmt.Errorf("token for user %v not found", userId)
}

func (s *Storage) DeleteToken(userId string) error {
	for i, token := range s.JWTokens {
		if token.UserID == userId {
			s.JWTokens = append(s.JWTokens[:i], s.JWTokens[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("token for user %v not found", userId)
}
