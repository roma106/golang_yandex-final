package main

import (
	"calculator_final/internal/app"
	"calculator_final/internal/logger"
)

func main() {
	app, err := app.New()
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		logger.Error("Failed to run application:" + err.Error())
	}
}
