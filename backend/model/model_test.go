package model

import (
	"context"
	"readinglistapp/initialisers"
	"readinglistapp/internal/data"
	"readinglistapp/internal/mocks"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var model = NewModel()

func TestInsert(t *testing.T) {
	mockCollection := &mocks.MockCollection{}
	bookCollection := &initialisers.BookCollection{Collection: mockCollection}

	mockInsertResult := &mongo.InsertOneResult{InsertedID: primitive.NewObjectID()}

	mockCollection.InsertOneFunc = func(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
		return mockInsertResult, nil
	}

	data := Input{
		Title:     "Unit Test",
		Published: 2022,
		Pages:     200,
		Genres:    []string{"Horror"},
		Rating:    2.2,
	}

	_, _, err := model.Insert(bookCollection, data)

	if err != nil {
		t.Errorf("got error %v, expected nil", err)
	}
}

func TestGet(t *testing.T) {
	mockCollection := &mocks.MockCollection{}
	bookCollection := &initialisers.BookCollection{Collection: mockCollection}

	bookID := "507f1f77bcf86cd799439011"

	expectedBook := &data.Book{
		ID:        bookID,
		Title:     "Test Book",
		Published: 2022,
		Pages:     123,
		Genres:    []string{"Fiction"},
		Rating:    4.5,
		Version:   1,
	}

	// Mocking the result for a successful find
	mockCollection.FindOneFunc = func(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
		return mongo.NewSingleResultFromDocument(expectedBook, nil, bson.DefaultRegistry)
	}

	_, err := model.Get(bookCollection, bookID)

	if err != nil {
		t.Errorf("got error %v, expected nil", err)
	}
}

func TestUpdateModel(t *testing.T) {
	// Create a mock instance of IBookCollection
	mockCollection := &mocks.MockCollection{}
	bookCollection := &initialisers.BookCollection{Collection: mockCollection}

	mockUpdateResult := &mongo.UpdateResult{}

	bookID := "507f1f77bcf86cd799439011"

	// Example data for update
	bookToUpdate := &data.Book{
		ID:      "507f1f77bcf86cd799439011",
		Title:   "Updated Book Title",
		Pages:   400,
		Genres:  []string{"Fantasy"},
		Rating:  4.2,
		Version: 2,
	}

	mockCollection.UpdateOneFunc = func(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
		// You can add assertions or custom logic here if needed
		return mockUpdateResult, nil
	}

	err := model.Update(bookCollection, bookID, bookToUpdate)

	if err != nil {
		t.Errorf("got error %v, expected nil", err)
	}
}

func TestDelete(t *testing.T) {
	// Create a mock instance of IBookCollection
	mockCollection := &mocks.MockCollection{}
	bookCollection := &initialisers.BookCollection{Collection: mockCollection}

	mockDeleteResult := &mongo.DeleteResult{}

	bookID := "507f1f77bcf86cd799439011"

	mockCollection.DeleteManyFunc = func(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
		return mockDeleteResult, nil
	}

	err := model.Delete(bookCollection, bookID)

	if err != nil {
		t.Errorf("got error %v, expected nil", err)
	}
}
