package expressions

import (
	"calculator_final/internal/databases"
	"calculator_final/internal/entities"
	"calculator_final/internal/logger"
	"calculator_final/internal/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func NewExpression(w http.ResponseWriter, r *http.Request, usersdb, exprsdb *sql.DB) {
	expr := entities.Expression{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to read request body. Error: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// парсим запрос
	reqExpr := struct {
		Username   string `json:"username"`
		Expression string `json:"expression"`
		Time       string `json:"time"`
	}{}
	err = json.Unmarshal(body, &reqExpr)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to unmarshal request body. Error: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// заполняем данные выражения
	expr.Expression = reqExpr.Expression
	expr.Status = "waiting"
	expr.StartingTime = utils.FormatTime(time.Now())
	reqtime, _ := strconv.Atoi(reqExpr.Time)
	expr.EndingTime = utils.FormatTime(time.Now().Add(time.Second * time.Duration(reqtime)))

	// получим ID пользователя
	user, err := databases.GetUserByUsername(usersdb, reqExpr.Username)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to get user by username from database. Error: %v", err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// отправляем в БД
	err = databases.AddExpression(exprsdb, user.ID, expr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
