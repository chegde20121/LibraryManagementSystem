package library

import (
	"testing"
)

func TestNewBook(t *testing.T) {
	// Positive test case
	book := NewBook("Title", "Author", 5)
	if book.Title != "Title" || book.Author != "Author" || book.Quantity != 5 {
		t.Errorf("NewBook() failed, expected: Title='Title', Author='Author', Quantity=5, got: %v", book)
	}
}

func TestBorrow(t *testing.T) {
	// Positive test case
	book := NewBook("Title", "Author", 5)
	err := book.Borrow()
	if err != nil {
		t.Errorf("Borrow() failed unexpectedly, error: %v", err)
	}
	if book.Quantity != 4 {
		t.Errorf("Borrow() failed, expected Quantity=4, got: %d", book.Quantity)
	}

	// Negative test case
	book.Quantity = 0 // Set quantity to 0 to simulate no available books
	err = book.Borrow()
	if err == nil {
		t.Errorf("Borrow() did not return expected error for no available books")
	}
}

func TestReturn(t *testing.T) {
	// Positive test case
	book := NewBook("Title", "Author", 5)
	book.Return()
	if book.Quantity != 6 {
		t.Errorf("Return() failed, expected Quantity=6, got: %d", book.Quantity)
	}
}
