package library

import (
	"fmt"
	"sync"
)

// Book represents a book in the library
type Book struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
	sync.Mutex
}

// NewBook creates a new book instance
func NewBook(title, author string, quantity int) *Book {
	return &Book{
		Title:    title,
		Author:   author,
		Quantity: quantity,
	}

}

// Borrow decrements the quantity of available books by 1, returns error if no books available
func (b *Book) Borrow() error {
	b.Lock()
	defer b.Unlock()
	if b.Quantity > 0 {
		b.Quantity--
		return nil
	}
	return fmt.Errorf("book not available")
}

// Return increments the quantity of available books by 1
func (b *Book) Return() {
	b.Lock()
	defer b.Unlock()
	b.Quantity++
}
