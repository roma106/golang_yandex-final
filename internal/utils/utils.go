package utils

import "net/http"

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	return nil
}

func RespondWithError(w http.ResponseWriter, code int, message string) error {
	return nil
}
