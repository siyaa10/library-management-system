package models

import "time"

type Book struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Author          string    `json:"author"`
	Category        string    `json:"category"`
	ISBN            string    `json:"isbn"`
	TotalCopies     int       `json:"totalCopies"`
	AvailableCopies int       `json:"availableCopies"`
	Description     *string   `json:"description"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// Check whether the book can be issued.
func (b *Book) IsAvailable() bool {
	return b.AvailableCopies > 0
}
