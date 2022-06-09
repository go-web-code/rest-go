package routes

import (
	"net/http"

	"github.com/go-web-code/rest-go/todo"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func SetupRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		r.Methods(route.Method).Path(route.Path).Name(route.Name).Handler(route.HandlerFunc)
	}

	return r
}

var routes = Routes{
	Route{"Index", "Get", "/", todo.ListTodoHandler},
}
