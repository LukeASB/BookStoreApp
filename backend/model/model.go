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

/* 💡LSB - Review - the DB initialisation should be in it's own class. Then we can dependency inject the DB/Collection into model. */
/* https://chatgpt.com/c/ff10255c-351b-4a40-8fea-bc4c1ccd6212 */
// var dbClient *initialisers.DB
// var collection *initialisers.BookCollection

/*
Sets up the DB Client.

Parameters:

	param1: pointer of DB client
*/
// func SetDBClient(client *initialisers.DB) {
// 	dbClient = client
// 	collection = initialisers.NewBookModel(dbClient)
// }

type Model struct {
}

func NewModel() *Model {
	return &Model{}
}

/*
Calls the DB to perform a create operation.

Parameters:

	param1: pointer of book data

Returns:

	return1: database id of inserted value
	return2: error
*/
func (m *Model) Insert(input Input) (interface{}, *data.Book, error) {
	data := &data.Book{
		ID:        "",
		Title:     input.Title,
		Published: input.Published,
		Pages:     input.Pages,
		Genres:    input.Genres,
		Rating:    input.Rating,
	}

	id, err := initialisers.Collection.Insert(data)
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
func (m *Model) GetAll() ([]*data.Book, error) {
	data, err := initialisers.Collection.GetAll()
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
func (m *Model) Get(id string) (*data.Book, error) {
	data, err := initialisers.Collection.Get(id)
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
func (m *Model) Update(id string, data *data.Book) error {
	err := initialisers.Collection.Update(data)
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
func (m *Model) Delete(id string) error {
	err := initialisers.Collection.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
