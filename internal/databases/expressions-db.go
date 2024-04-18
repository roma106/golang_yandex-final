package databases

import (
	"calculator_final/internal/entities"
	"calculator_final/internal/logger"
	"calculator_final/internal/utils"
	"database/sql"
	"fmt"
	"time"
)

func CreateExpressionsTable(db *sql.DB, userID string) error {
	_, err := db.Exec(fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS "Expressions-User%s"
	(
		"ID" bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
		"Expression" character varying NOT NULL,
		"Status" character varying NOT NULL,
		"Result" bigint,
		"StartingTime" timestamp without time zone NOT NULL,
		"EndingTime" timestamp without time zone,
		CONSTRAINT "calculator-expressions%s_pkey" PRIMARY KEY ("ID")
	)`,
		userID, userID))
	if err != nil {
		logger.Error("Failed to create expressions table for user " + userID + ": " + err.Error())
		return err
	}
	logger.Info(fmt.Sprintf("Table 'Expressions-User%s' created for user %s", userID, userID))
	return nil
}

func GetExpressions(db *sql.DB, userID string) ([]entities.Expression, error) {
	rows, err := db.Query(fmt.Sprintf(`SELECT * FROM "Expressions-User%s"`, userID))
	if err != nil {
		logger.Error("Failed to select expressions from table for user " + userID + ": " + err.Error())
		return nil, err
	}
	defer rows.Close()
	exprs := []entities.Expression{}
	for rows.Next() {
		expr := entities.Expression{}
		result := new(interface{})
		err = rows.Scan(&expr.ID, &expr.Expression, &expr.Status, &result, &expr.StartingTime, &expr.EndingTime)
		if err != nil {
			logger.Error("Failed to scan expressions from table for user " + userID + ": " + err.Error())
			continue
		}
		expr.StartingTime = utils.FormatTime(expr.StartingTime)
		// считаем выражение и проверяем статус выполнения
		expr.Status = UpdateExpressionStatus(db, userID, expr)
		res, err := utils.CalculateExpression(expr.Expression)
		if err != nil {
			return nil, err
		}
		expr.Result = res
		exprs = append(exprs, expr)
	}
	return exprs, nil
}

func AddExpression(db *sql.DB, userID string, expr entities.Expression) error {
	_, err := db.Exec(fmt.Sprintf(`
	INSERT INTO "Expressions-User%s" ("Expression", "Status", "StartingTime", "EndingTime") VALUES ($1, $2, $3, $4)`, userID),
		expr.Expression, "waiting", expr.StartingTime, expr.EndingTime)
	if err != nil {
		logger.Error("Failed to add expression to table for user " + userID + ": " + err.Error())
		return err
	}
	logger.Info(fmt.Sprintf("Expression %s added to table for user %s", expr.Expression, userID))
	return nil
}

func UpdateExpressionStatus(db *sql.DB, userID string, expr entities.Expression) string {
	if expr.Status == "done" {
		return "done"
	}
	if !expr.EndingTime.After(utils.FormatTime(time.Now())) {
		updateStmt, err := db.Prepare(fmt.Sprintf(`UPDATE "Expressions-User%s" SET "Result" = $1, "Status" = $2 WHERE "ID" = $3`, userID))
		if err != nil {
			fmt.Println("Ошибка при подготовке запроса обновления:", err)
			return "failed"
		}
		defer updateStmt.Close()
		res, err := utils.CalculateExpression(expr.Expression)
		if err != nil {
			return "failed"
		}
		_, err = updateStmt.Exec(res, "done", expr.ID)
		if err != nil {
			logger.Error("Failed to update expression status in database: " + err.Error())
			return "failed"
		}

		logger.Info(fmt.Sprintf("Expression status %s updated in table for user %s", expr.Expression, userID))
		return "done"
	} else {
		return "waiting"
	}
}
