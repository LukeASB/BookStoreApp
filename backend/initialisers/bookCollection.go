package initialisers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"readinglistapp/internal/data"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IBookCollection interface {
	Create(book *data.Book) (interface{}, error)
	Delete(id string) error
	Get(id string) (*data.Book, error)
	GetAll() ([]*data.Book, error)
	Update(book *data.Book) error
}

type ICollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type BookCollection struct {
	Collection ICollection
}

func NewBookCollection(client *DB) *BookCollection {
	return NewBookModel(client)
}

/*
NewBookModel creates a new instance of BookCollection, which represents a collection of books in the MongoDB database.
It takes a pointer to a DB struct as input, representing the database connection, and returns a pointer to a BookCollection.
The BookCollection is initialized with the collection named "books" in the "readinglist" database.

Parameters:

param1: pointer DB

Returns:

return1: pointer BookCollection
*/
func NewBookModel(db *DB) *BookCollection {
	return &BookCollection{Collection: db.client.Database("readinglist").Collection("books")}
}

/*
Insert inserts a new book into the BookCollection.
It takes a pointer to a Book struct as input and returns the ID of the inserted document and an error.
If the CreatedAt timestamp is not set in the input book, it sets the current time as the CreatedAt timestamp.

Parameters:
param1: pointer Book

Returns:
return1: interface{}, ID of the inserted document
return2: error
*/
func (bc *BookCollection) Create(book *data.Book) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Set CreatedAt timestamp if not already set
	if book.CreatedAt.IsZero() {
		book.CreatedAt = time.Now()
	}

	data := data.BookData{
		ID:        primitive.NewObjectID(),
		CreatedAt: book.CreatedAt,
		Title:     book.Title,
		Published: book.Published,
		Pages:     book.Pages,
		Genres:    book.Genres,
		Rating:    book.Rating,
		Version:   book.Version,
	}

	result, err := bc.Collection.InsertOne(ctx, data)

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

/*
Get retrieves a book from the BookCollection by its ID.
It takes a string representing the ID of the book as input and returns a pointer to the retrieved Book struct and an error.
If the document with the specified ID is not found, it returns a "record not found" error.

Parameters:
param1: string, ID of the book

Returns:
return1: pointer Book
return2: error
*/
func (bc *BookCollection) Get(id string) (*data.Book, error) {
	objID, err := parseToObjectID(id)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: objID}}

	var result data.Book

	err = bc.Collection.FindOne(ctx, filter).Decode(&result)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}

	return &result, nil
}

/*
GetAll retrieves all books from the BookCollection.
It returns a slice of pointers to Book structs and an error.

Returns:
return1: []*Book, slice of pointers to Book structs
return2: error
*/
func (bc *BookCollection) GetAll() ([]*data.Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cur, err := bc.Collection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	var results []*data.Book

	for cur.Next(context.TODO()) {
		var elem data.BookData
		err := cur.Decode(&elem)

		if err != nil {
			log.Fatal(err)
		}

		idStr := elem.ID.Hex()

		book := data.Book{
			ID:        idStr,
			CreatedAt: elem.CreatedAt,
			Title:     elem.Title,
			Published: elem.Published,
			Pages:     elem.Pages,
			Genres:    elem.Genres,
			Rating:    elem.Rating,
			Version:   elem.Version,
		}

		results = append(results, &book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return results, nil
}

/*
Update updates a book in the BookCollection.
It takes a pointer to a Book struct as input and updates the document with the corresponding ID in the collection.
It returns an error.

Parameters:
param1: pointer Book

Returns:
return1: error
*/
func (bc *BookCollection) Update(book *data.Book) error {
	objID, err := parseToObjectID(book.ID)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create a filter to find the document by its ID
	filter := bson.D{{Key: "_id", Value: objID}}

	// Create an update with the changes to apply
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "title", Value: book.Title},
			{Key: "published", Value: book.Published},
			{Key: "pages", Value: book.Pages},
			{Key: "genres", Value: book.Genres},
			{Key: "rating", Value: book.Rating},
			{Key: "version", Value: book.Version},
		}},
	}

	// Perform the update operation
	result, err := bc.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("\nThe document has been updated. ModifiedCount: %v, UpdatedCount: %v, UpdatedID: %v", result.ModifiedCount, result.UpsertedCount, result.UpsertedID)

	return nil
}

/*
Delete removes a book from the BookCollection by its ID.
It takes a string representing the ID of the book as input and returns an error.

Parameters:
param1: string, ID of the book

Returns:
return1: error
*/
func (bc *BookCollection) Delete(id string) error {
	objID, err := parseToObjectID(id)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Create a filter to find the document by its ID
	filter := bson.D{{Key: "_id", Value: objID}}

	_, err = bc.Collection.DeleteMany(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}

/*
parseToObjectID converts a string representation of ObjectID to a primitive.ObjectID object.
It takes a string representing the ObjectID as input and returns the corresponding primitive.ObjectID object and an error.

Parameters:
param1: string, representation of ObjectID

Returns:
return1: primitive.ObjectID, parsed ObjectID
return2: error
*/
func parseToObjectID(id string) (primitive.ObjectID, error) {
	if len(id) <= 0 {
		return primitive.NilObjectID, errors.New("record could not be found")
	}
	// Parse the string representation of ObjectID into a primitive.ObjectID object
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return objID, nil
}
