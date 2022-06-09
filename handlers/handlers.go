package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-web-code/rest-go/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("====> Index Handler")
	todo := models.Todo{}
	json.NewEncoder(w).Encode(todo)
}
