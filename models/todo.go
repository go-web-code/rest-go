package models

import "time"

type Todo struct {
	Title     string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}
