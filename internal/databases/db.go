package databases

import (
	"calculator_final/internal/logger"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitUsers(db *sql.DB) error {
	logger.Info("Initializing users table...")
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS "calculator-users"
	(
		"ID" bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
		"Username" character varying NOT NULL,
		"Password" character varying NOT NULL,
		"LastLoggedIn" timestamp without time zone,
		CONSTRAINT "calculator-users_pkey" PRIMARY KEY ("ID")
	)
	`)
	if err != nil {
		logger.Error("Failed to create table: " + err.Error())
		return err
	}
	logger.Info("Users table created!")
	return nil
}

func ConnectToDB(name string) *sql.DB {
	logger.Info(fmt.Sprintf("Connecting to database %s...", name))
	db, err := sql.Open("postgres", fmt.Sprintf("host=postgres port=5432 user=postgres password=ri106rom dbname=%s sslmode=disable", name))
	if err != nil {
		logger.Error("Failed to connect to database: " + err.Error())
		return nil
	}
	err = db.Ping()
	if err != nil {
		logger.Error("Failed to connect to database: " + err.Error())
		return nil
	}
	logger.Info("Database " + name + " connected!")
	return db
}
