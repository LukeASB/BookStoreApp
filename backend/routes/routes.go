package routes

import (
	"net/http"
	controller "readinglistapp/controller"
	"readinglistapp/model"
	view "readinglistapp/view"

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
	v := view.NewView()
	m := model.NewModel()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server handler with the /static/ route prefix
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controller.Home(w, r, v, m)
	})
	router.HandleFunc("/book/view", func(w http.ResponseWriter, r *http.Request) {
		controller.BookView(w, r, v, m)
	})
	router.HandleFunc("/book/create", func(w http.ResponseWriter, r *http.Request) {
		controller.BookCreate(w, r, v)
	})

	router.HandleFunc("/v1/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		controller.HealthCheck(w, r, v)
	})

	router.HandleFunc("/v1/books", func(w http.ResponseWriter, r *http.Request) {
		controller.GetBooksHandler(w, r, v, m)
	}).Methods(http.MethodGet)

	router.HandleFunc("/v1/books", func(w http.ResponseWriter, r *http.Request) {
		controller.CreateBooksHandler(w, r, v, m)
	}).Methods(http.MethodPost)

	router.HandleFunc("/v1/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.GetBook(w, r, v, m)
	}).Methods(http.MethodGet)

	router.HandleFunc("/v1/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.UpdateBook(w, r, v, m)
	}).Methods(http.MethodPut)

	router.HandleFunc("/v1/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		controller.DeleteBook(w, r, v, m)
	}).Methods(http.MethodDelete)
}
