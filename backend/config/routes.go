package config

import (
	"net/http"
	"readinglistapp/initialisers"
	"readinglistapp/model"
	"readinglistapp/routes"
	"readinglistapp/view"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*
SetUpRouter creates and configures a new HTTP router using mux.Router.
It sets up routes defined in the routes package, enables handling of trailing slashes,
and applies CORS (Cross-Origin Resource Sharing) middleware to allow requests from any origin.

Returns:

	return1: http.handler that can be used to serve HTTP requests.
*/
func SetUpRouter(v *view.View, m *model.Model, db *initialisers.BookCollection) http.Handler {
	muxRouter := mux.NewRouter()

	routes.SetUpRoutes(muxRouter, v, m, db)

	muxRouter.StrictSlash(false)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow requests from your React app's origin
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Allow sending cookies and credentials
	}).Handler(muxRouter)

	return corsHandler
}
