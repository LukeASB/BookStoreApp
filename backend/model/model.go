package model

import (
	"readinglistapp/initialisers"
	"readinglistapp/internal/data"
)

type Input struct {
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     int      `json:"pages"`
	Genres    []string `json:"genres"`
	Rating    float64  `json:"rating"`
}

var dbClient *initialisers.DB
var collection *initialisers.BookCollection

/*
Sets up the DB Client.

Parameters:

	param1: pointer of DB client
*/
func SetDBClient(client *initialisers.DB) {
	dbClient = client
	collection = initialisers.NewBookModel(dbClient)
}

/*
Calls the DB to perform a create operation.

Parameters:

	param1: pointer of book data

Returns:

	return1: database id of inserted value
	return2: error
*/
func Insert(input Input) (interface{}, *data.Book, error) {
	data := &data.Book{
		ID:        "",
		Title:     input.Title,
		Published: input.Published,
		Pages:     input.Pages,
		Genres:    input.Genres,
		Rating:    input.Rating,
	}

	id, err := collection.Insert(data)
	if err != nil {
		return nil, nil, err
	}
	return id, data, nil
}

/*
Calls the DB to perform a retrieve operation.

Returns:

	return1: slice of a pointer of books
	return2: error
*/
func GetAll() ([]*data.Book, error) {
	data, err := collection.GetAll()
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
Calls the DB to perform a retrieve operation with a given id.

Parameters:

	param1: id string

Returns:

	return1: slice of a pointer of books
	return2: error
*/
func Get(id string) (*data.Book, error) {
	data, err := collection.Get(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

/*
Calls the DB to perform an update operation with a given id and data.

Parameters:

	param1: id string
	param2: pointer of book data

Returns:

	return1: slice of a pointer of books
	return2: error
*/
func Update(id string, data *data.Book) error {
	err := collection.Update(data)
	if err != nil {
		return err
	}
	return nil
}

/*
Calls the DB to perform a delete operation with a given id.

Parameters:

	param1: id string

Returns:

	return1: error
*/
func Delete(id string) error {
	err := collection.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
