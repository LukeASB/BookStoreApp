package data

import (
	"time"
)

type Book struct {
	ID        string    `json:"_id" bson:"_id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Published int       `json:"published,omitempty"`
	Pages     int       `json:"pages,omitempty"`
	Genres    []string  `json:"genres,omitempty"`
	Rating    float64   `json:"rating,omitempty"`
	Version   int32     `json:"version,omitempty"`
}
