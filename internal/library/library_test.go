package library

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetLibraryInstance(t *testing.T) {
	tempDir := t.TempDir()
	viper.Set("DATA_FILEPATH", tempDir)

	// Create a test JSON file
	booksData := `[{"Title":"Book1","Author":"Author1","Quantity":3},{"Title":"Book2","Author":"Author2","Quantity":2}]`
	filePath := filepath.Join(tempDir, "books.json")
	file, err := os.Create(filePath)
	assert.NoError(t, err)
	defer file.Close()
	_, err = file.WriteString(booksData)
	assert.NoError(t, err)
	instance1 := GetLibraryInstance()
	instance2 := GetLibraryInstance()

	assert.NotNil(t, instance1)
	assert.NotNil(t, instance2)
	assert.Equal(t, instance1, instance2)
}

func TestLoadBooks(t *testing.T) {
	// Mock data directory
	tempDir := t.TempDir()
	viper.Set("DATA_FILEPATH", tempDir)

	// Create a test JSON file
	booksData := `[{"Title":"Book1","Author":"Author1","Quantity":3},{"Title":"Book2","Author":"Author2","Quantity":2}]`
	filePath := filepath.Join(tempDir, "books.json")
	file, err := os.Create(filePath)
	assert.NoError(t, err)
	defer file.Close()
	_, err = file.WriteString(booksData)
	assert.NoError(t, err)

	// Test loading books from the JSON file
	library, err := loadBooks()
	assert.NoError(t, err)
	assert.NotNil(t, library)
	assert.Len(t, library, 2)

	// Test error case (file not found)
	viper.Set("DATA_FILEPATH", "/nonexistent")
	_, err = loadBooks()
	assert.Error(t, err)
}

func TestGetAllBooks(t *testing.T) {
	library := &Library{
		books: map[string]*Book{
			"Book1": {Title: "Book1", Author: "Author1", Quantity: 3},
			"Book2": {Title: "Book2", Author: "Author2", Quantity: 2},
		},
	}

	books := library.GetAllBooks()
	assert.NotNil(t, books)
	assert.Len(t, books, 2)
}

func TestAddBookByTitle(t *testing.T) {
	library := &Library{
		books: make(map[string]*Book),
	}

	book := &Book{Title: "Book1", Author: "Author1", Quantity: 3}
	library.AddBookByTitle(book)

	assert.NotNil(t, library.books["Book1"])
	assert.Equal(t, 3, library.books["Book1"].Quantity)

	// Test adding existing book
	library.AddBookByTitle(book)
	assert.Equal(t, 4, library.books["Book1"].Quantity)
}

func TestGetBookByTitle(t *testing.T) {
	library := &Library{
		books: map[string]*Book{
			"Book1": {Title: "Book1", Author: "Author1", Quantity: 3},
		},
	}

	book := library.GetBookByTitle("Book1")
	assert.NotNil(t, book)
	assert.Equal(t, "Book1", book.Title)

	book = library.GetBookByTitle("NonexistentBook")
	assert.Nil(t, book)
}
