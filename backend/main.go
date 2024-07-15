package main

import (
	"fmt"
	"net/http"
	"os"
	"readinglistapp/config"
	"readinglistapp/initialisers"
	"readinglistapp/internal"
)

var (
	app     internal.App
	DB      *initialisers.DB
	SiteURL string
)

/*
init is called before the main function and initializes the application by loading environment variables
and connecting to the database.
If an error occurs during database connection, it panics.
*/
func init() {
	initialisers.LoadEnvVariables()
	DB = app.NewDB()
}

/*
main is the entry point of the application. It sets up the router, initializes the database client,
and starts the server to listen on the specified port.
*/
func main() {
	port := os.Getenv("PORT")

	app = internal.App{
		View:  app.NewView(),
		Model: app.NewModel(),
		DB:    DB,
	}

	router := config.SetUpRouter(app)

	defer cleanup(DB.Close)

	fmt.Printf("\nListening on port: %s", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}

/*
Closes DB connection
*/
func cleanup(disconnect func()) {
	fmt.Println("\nExecuting Clean Up...")
	defer disconnect()
}
