package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-web-code/rest-go/common"
)

func ListTodoHandler(w http.ResponseWriter, r *http.Request) {

	todo := Todo{}

	json.NewDecoder(r.Body).Decode(&todo)
	log.Println(todo)
	todo = Todo{Title: "task work 1", Completed: false, Due: common.JSONTime(time.Now().AddDate(0, 0, 2))}
	json.NewEncoder(w).Encode(todo)

}
