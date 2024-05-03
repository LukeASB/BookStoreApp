package routes

import (
	"net/http"
	controller "readinglistapp/controller"

	"github.com/gorilla/mux"
)

/*
SetUpRoutes configures the router with appropriate handlers for different endpoints.
It serves static files for UI assets, defines routes for home page, book view, creation,
health check endpoint, and CRUD operations for books under /v1/books endpoint.

Parameters:

	param1: gorilla mux router
*/
func SetUpRoutes(router *mux.Router) {
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server handler with the /static/ route prefix
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	router.HandleFunc("/", controller.Home)
	router.HandleFunc("/book/view", controller.BookView)
	router.HandleFunc("/book/create", controller.BookCreate)

	router.HandleFunc("/v1/healthcheck", controller.HealthCheck)

	router.HandleFunc("/v1/books", controller.GetBooksHandler).Methods(http.MethodGet)
	router.HandleFunc("/v1/books", controller.CreateBooksHandler).Methods(http.MethodPost)
	router.HandleFunc("/v1/books/{id}", controller.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/v1/books/{id}", controller.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/v1/books/{id}", controller.DeleteBook).Methods(http.MethodDelete)
}
