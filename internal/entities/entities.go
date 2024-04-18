package entities

import "time"

type JWT struct {
}

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	LastLogged time.Time
}

type Expression struct {
	ID           string    `json:"id"`
	Expression   string    `json:"expression"`
	Status       string    `json:"status"`
	Result       float64   `json:"result"`
	StartingTime time.Time `json:"starting_time"`
	EndingTime   time.Time `json:"ending_time"`
}
