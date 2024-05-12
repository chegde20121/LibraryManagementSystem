package user

import (
	"bufio"
	"strings"
	"testing"

	"github.com/chegde20121/LibraryManagementSystem/internal/library"
	"github.com/spf13/viper"
)

func init() {
	viper.Set("DATA_FILEPATH", "../../internal/library/")
}
func TestLogin(t *testing.T) {
	testCases := []struct {
		username string
		password string
		expected bool
	}{
		{"admin", "admin123", true},        // Valid credentials
		{"admin", "wrongpass", false},      // Invalid password
		{"nonexistent", "password", false}, // Nonexistent user
	}

	for _, tc := range testCases {
		u, err := Login(tc.username, tc.password)
		if tc.expected && (err != nil || u == nil) {
			t.Errorf("Expected login to succeed for user '%s', but it failed", tc.username)
		} else if !tc.expected && (err == nil || u != nil) {
			t.Errorf("Expected login to fail for user '%s', but it succeeded", tc.username)
		}
	}
}

func TestAddNewUser(t *testing.T) {
	// Create a mocked reader
	reader := bufio.NewReader(strings.NewReader("2\n"))
	u := &User{UserName: "test", Password: "1234", Role: ADMIN_ROLE}
	u.AddNewUser(reader)
	// Verify the user has been added
	if len(users) != 4 {
		t.Error("Expected 4 users after adding a new one, but got:", len(users))
	}

}

func TestViewAllBooks(t *testing.T) {
	// Create some mock books in the library
	lib := library.GetLibraryInstance()
	lib.AddBookByTitle(&library.Book{Title: "Book1", Author: "Author1", Quantity: 2})
	lib.AddBookByTitle(&library.Book{Title: "Book2", Author: "Author2", Quantity: 1})

	u := &User{}
	u.ViewAllBooks()
}

func TestBorrowBook(t *testing.T) {
	// Create a mock book in the library
	lib := library.GetLibraryInstance()
	lib.AddBookByTitle(&library.Book{Title: "Book1", Author: "Author1", Quantity: 2})

	u := &User{}

	// Borrow a book
	err := u.BorrowBook("Book1")
	if err != nil {
		t.Error("Expected no error when borrowing a book, but got:", err)
	}

	// Borrow a non-existent book
	err = u.BorrowBook("NonExistentBook")
	if err == nil {
		t.Error("Expected an error when borrowing a non-existent book, but got nil")
	}
}

func TestReturnBook(t *testing.T) {
	// Create a mock book in the library
	lib := library.GetLibraryInstance()
	lib.AddBookByTitle(&library.Book{Title: "Book1", Author: "Author1", Quantity: 1})

	u := &User{}
	u.BorrowBook("Book1")

	// Return a book
	err := u.ReturnBook("Book1")
	if err != nil {
		t.Error("Expected no error when returning a book, but got:", err)
	}

	// Return a non-existent borrowed book
	err = u.ReturnBook("NonBorrowedBook")
	if err == nil {
		t.Error("Expected an error when returning a non-borrowed book, but got nil")
	}
}

func TestDisplayBooksBorrowed(t *testing.T) {
	u := &User{}
	// Display borrowed books when none are borrowed
	err := u.DisplayBooksBorrowed()
	if err == nil {
		t.Error("Expected an error when displaying borrowed books with none borrowed, but got nil")
	}

	// Borrow a book
	lib := library.GetLibraryInstance()
	lib.AddBookByTitle(&library.Book{Title: "Book1", Author: "Author1", Quantity: 1})
	u.BorrowBook("Book1")

	// Display borrowed books
	err = u.DisplayBooksBorrowed()
	if err != nil {
		t.Error("Expected no error when displaying borrowed books, but got:", err)
	}
}

func TestSearchBook(t *testing.T) {
	// Create some mock books in the library
	lib := library.GetLibraryInstance()
	lib.AddBookByTitle(&library.Book{Title: "Book1", Author: "Author1", Quantity: 2})
	lib.AddBookByTitle(&library.Book{Title: "Book2", Author: "Author2", Quantity: 1})

	u := &User{}

	// Search for an existing book by title
	err := u.SearchBook(lib, bufio.NewReader(strings.NewReader("Book1\n")))
	if err != nil {
		t.Error("Expected no error when searching for an existing book, but got:", err)
	}

	// Search for a non-existent book by title
	err = u.SearchBook(lib, bufio.NewReader(strings.NewReader("NonExistentBook\n")))
	if err == nil {
		t.Error("Expected an error when searching for a non-existent book by title, but got nil")
	}

	// Search for an existing book by author
	err = u.SearchBook(lib, bufio.NewReader(strings.NewReader("Author2\n")))
	if err != nil {
		t.Error("Expected no error when searching for an existing book by author, but got:", err)
	}

	// Search for a non-existent book by author
	err = u.SearchBook(lib, bufio.NewReader(strings.NewReader("NonExistentAuthor\n")))
	if err == nil {
		t.Error("Expected an error when searching for a non-existent book by author, but got nil")
	}
}
