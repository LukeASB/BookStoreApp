package main

import (
	"fmt"
	"net/http"
	"os"
	"readinglistapp/config"
	"readinglistapp/initialisers"
	"readinglistapp/model"
	"readinglistapp/view"
)

var (
	dbClient *initialisers.DB
	SiteURL  string
)

/*
init is called before the main function and initializes the application by loading environment variables
and connecting to the database.
If an error occurs during database connection, it panics.
*/
func init() {
	var err error
	initialisers.LoadEnvVariables()
	dbClient, err = initialisers.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	initialisers.SetDBClient(dbClient)
}

/*
main is the entry point of the application. It sets up the router, initializes the database client,
and starts the server to listen on the specified port.
*/
func main() {
	port := os.Getenv("PORT")

	db := initialisers.Collection
	v := view.NewView()
	m := model.NewModel()

	router := config.SetUpRouter(v, m, db)

	defer cleanup(dbClient.Close)

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
