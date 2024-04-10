package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"jwt_lesson/internal/entities"
	"jwt_lesson/internal/logger"
	"net/http"
)

func LogIn(w http.ResponseWriter, r *http.Request) {
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
	logger.Info(fmt.Sprintf("user %v successfully logged in", user.Username))
}
