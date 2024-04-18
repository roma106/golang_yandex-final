package utils

import (
	"calculator_final/internal/logger"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Knetic/govaluate"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Write(body)
	return nil
}

func FormatTime(t time.Time) time.Time {
	formattedTime := t.Format("2006-01-02 15:04:05")
	newTime, _ := time.Parse("2006-01-02 15:04:05", formattedTime)
	return newTime
}

func CalculateExpression(expr string) (float64, error) {
	expression, _ := govaluate.NewEvaluableExpression(expr)
	result, err := expression.Evaluate(nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to calculate expression. Error: %v", err))
		return 0, fmt.Errorf("ошибка вычисления: %v", err)
	}
	return result.(float64), nil
}
