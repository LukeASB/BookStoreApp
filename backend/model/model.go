package model

import (
	"readinglistapp/initialisers"
	"readinglistapp/internal/data"
)

type IModelNew interface {
	NewModel() *Model
}

type IModelFuncs interface {
	Delete(db initialisers.IBookCollection, id string) error
	Get(db initialisers.IBookCollection, id string) (*data.Book, error)
	GetAll(db initialisers.IBookCollection) ([]*data.Book, error)
	Insert(db initialisers.IBookCollection, input Input) (interface{}, *data.Book, error)
	Update(db initialisers.IBookCollection, id string, data *data.Book) error
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
func (m *Model) Insert(db initialisers.IBookCollection, input Input) (interface{}, *data.Book, error) {
	data := &data.Book{
		ID:        "",
		Title:     input.Title,
		Published: input.Published,
		Pages:     input.Pages,
		Genres:    input.Genres,
		Rating:    input.Rating,
	}

	id, err := db.Create(data)
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
func (m *Model) GetAll(db initialisers.IBookCollection) ([]*data.Book, error) {
	data, err := db.GetAll()
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
func (m *Model) Get(db initialisers.IBookCollection, id string) (*data.Book, error) {
	data, err := db.Get(id)
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
func (m *Model) Update(db initialisers.IBookCollection, id string, data *data.Book) error {
	err := db.Update(data)
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
func (m *Model) Delete(db initialisers.IBookCollection, id string) error {
	err := db.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
