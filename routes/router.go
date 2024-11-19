package routes

import (
	"e-book_webapp/handlers"
	"net/http"
)

func SetupRouter() *http.ServeMux {

	router := http.NewServeMux()
	router.HandleFunc("/", handlers.HomeHandler)

	return router
}
