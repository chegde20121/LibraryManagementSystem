package library

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var (
	instance *Library
	once     sync.Once
)

type Library struct {
	books map[string]*Book
	sync.Mutex
}

func GetLibraryInstance() *Library {
	once.Do(func() {
		loadedBooks, _ := loadBooks()
		instance = &Library{
			books: loadedBooks,
		}
	})
	return instance
}

func loadBooks() (map[string]*Book, error) {
	fmt.Println(viper.Get("DATA_FILEPATH").(string))
	path := filepath.Join(viper.Get("DATA_FILEPATH").(string), "books.json")
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	// Decode the JSON data into a slice of Book objects
	var books []*Book
	booksLoaded := make(map[string]*Book)

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&books); err != nil {
		return nil, err
	}
	// Add the books to the library
	for _, book := range books {
		if existingbook, ok := booksLoaded[book.Title]; ok {
			existingbook.Quantity++
		} else {
			booksLoaded[book.Title] = book
		}
	}
	return booksLoaded, nil
}

func (l *Library) GetAllBooks() map[string]*Book {
	return l.books
}

// AddBook adds a book to the library
func (l *Library) AddBookByTitle(book *Book) {
	l.Lock()
	defer l.Unlock()
	if existingBook, ok := l.books[book.Title]; ok {
		existingBook.Quantity++
	} else {
		l.books[book.Title] = book
	}
}
func (l *Library) GetBookByTitle(title string) *Book {
	return l.books[title]
}

func (lib *Library) GetBooksByAuthor(author string) []*Book {
	var booksByAuthor []*Book
	for _, book := range lib.books {
		if book.Author == author {
			booksByAuthor = append(booksByAuthor, book)
		}
	}
	return booksByAuthor
}
