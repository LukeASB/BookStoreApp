package initialisers

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDB interface {
	NewDB() *DB
}

type DB struct {
	client *mongo.Client
}

func NewDB() (*DB, error) {
	dbClient, err := connectToDatabase()

	if err != nil {
		return nil, err
	}

	return dbClient, nil
}

/*
ConnectToDatabase establishes a connection to the MongoDB database using the provided environment
variable DB_URL.
It creates a new MongoDB client and attempts to ping the server to confirm the
connection.
If successful, it returns a pointer to the DB struct containing the client.
If there is an error during connection or ping, it returns nil and the error.

Returns:

	return1: pointer DB
	return2: error
*/
func connectToDatabase() (*DB, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(os.Getenv("DB_URL")).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, err
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	_, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return &DB{client: client}, nil
}

/*
Close closes the connection to the MongoDB database. It calls the Disconnect method on the MongoDB
client associated with the DB struct. If there is an error during disconnection, it panics with the error.
Otherwise, it prints a message indicating successful disconnection.

Returns:

return1: pointer DB
*/
func (db *DB) Close() {
	if err := db.client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}

	fmt.Println("\nSuccessfully disconnected from MongoDB")
}
