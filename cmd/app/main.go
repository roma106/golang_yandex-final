package main

import (
	"jwt_lesson/internal/app"
	"jwt_lesson/internal/logger"
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
