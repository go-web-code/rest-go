package todo

import "github.com/go-web-code/rest-go/common"

type Todo struct {
	Title     string          `json:"title"`
	Completed bool            `json:"completed"`
	Due       common.JSONTime `json:"due"`
}
