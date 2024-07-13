package initialisers

import (
	"testing"
)

// Add mock crud operations that don't call the DB.

func TestUpdate(t *testing.T) {
	// mockCollection := &mocks.MockCollection{}
	// bookCollection := &BookCollection{Collection: mockCollection}

	// book := &data.BookData{
	// 	ID:        "id",
	// 	CreatedAt: time.Now(),
	// 	Title:     "Unit Test",
	// 	Published: 2023,
	// 	Pages:     123,
	// 	Genres:    []string{"Horror"},
	// 	Rating:    4.5,
	// 	Version:   2,
	// }
	// type Book struct {
	// 	ID        string    `json:"_id" bson:"_id"`
	// 	CreatedAt time.Time `json:"createdAt"`
	// 	Title     string    `json:"title"`
	// 	Published int       `json:"published,omitempty"`
	// 	Pages     int       `json:"pages,omitempty"`
	// 	Genres    []string  `json:"genres,omitempty"`
	// 	Rating    float64   `json:"rating,omitempty"`
	// 	Version   int32     `json:"version,omitempty"`
	// }
}
