package user

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/chegde20121/LibraryManagementSystem/internal/library"
)

type User struct {
	UserName string                   `json:"username"`
	Password string                   `json:"password"`
	Role     string                   `json:"role"`
	Borrowed map[string]*library.Book `json:"booksBorrowed"`
}

const (
	ADMIN_ROLE  = "ADMIN"
	READER_ROLE = "READER"
)

var users = []*User{
	{UserName: "admin", Password: "admin123", Role: ADMIN_ROLE},
	{UserName: "user1", Password: "password1", Role: READER_ROLE},
	{UserName: "user2", Password: "password2", Role: READER_ROLE},
}

func Login(username string, password string) (*User, error) {

	for _, user := range users {
		if user.UserName == username && user.Password == password {
			return user, nil
		}
	}
	return nil, fmt.Errorf("invalid username or password")
}
func (u *User) AddNewUser(reader *bufio.Reader) {
	var username, password string
	var role int
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter new password: ")
	fmt.Scanln(&password)
	fmt.Println("Chose Role ")
	fmt.Println("1. ADMIN_ROLE")
	fmt.Println("2. READER_ROLE")
	data, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("Invalid choice. Please try again.")
		return
	}
	role, err = strconv.Atoi(string(data))
	if err != nil {
		fmt.Println("Invalid choice. Please try again.")
		return
	}
	var assigned_role string
	switch role {
	case 1:
		{
			assigned_role = ADMIN_ROLE
		}
	case 2:
		{
			assigned_role = READER_ROLE
		}
	default:
		{
			fmt.Println("you have entered wrong choice")
			return
		}
	}
	user := User{
		UserName: username,
		Password: password,
		Role:     assigned_role,
		Borrowed: make(map[string]*library.Book),
	}
	users = append(users, &user)
	fmt.Println("successfully added new user")
}

func (u *User) ViewAllBooks() {
	//get library instance
	lib := library.GetLibraryInstance()
	fmt.Println("Available Books")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', 0)
	// Print table headers
	fmt.Fprintf(w, "Title\tAuthor\tQuantity\n")
	for _, book := range lib.GetAllBooks() {
		fmt.Fprintf(w, "%s\t%s\t%d\n", book.Title, book.Author, book.Quantity)
	}
	w.Flush()
}

func (u *User) BorrowBook(keyword string) error {
	lib := library.GetLibraryInstance()
	book := lib.GetBookByTitle(keyword)
	if book == nil || book.Quantity == 0 {
		return fmt.Errorf("book %s not present", keyword)
	}
	if u.Borrowed == nil {
		u.Borrowed = make(map[string]*library.Book)
	}
	if ownedBook, ok := u.Borrowed[book.Title]; ok {

		ownedBook.Quantity++

	} else {
		u.Borrowed[book.Title] = &library.Book{
			Title:    book.Title,
			Author:   book.Author,
			Quantity: 1,
		}
	}
	book.Borrow()
	return nil
}

func (u *User) ReturnBook(keyword string) error {
	lib := library.GetLibraryInstance()
	if len(u.Borrowed) > 0 {
		if borrowedBook, ok := u.Borrowed[keyword]; ok {
			book := lib.GetBookByTitle(borrowedBook.Title)
			borrowedBook.Quantity--
			if borrowedBook.Quantity == 0 {
				delete(u.Borrowed, borrowedBook.Title)
			}
			book.Return()
			return nil
		}
	}
	return fmt.Errorf("you haven't borrowed any book")
}

func (u *User) DisplayBooksBorrowed() error {
	if len(u.Borrowed) < 1 {
		return fmt.Errorf("you haven't borrowed any books yet")
	}
	fmt.Println("Borrowed Books:")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 8, ' ', 0)
	// Print table headers
	fmt.Fprintf(w, "Title\tAuthor\tQuantity\n")
	for _, book := range u.Borrowed {
		fmt.Fprintf(w, "%s\t%s\t%d\n", book.Title, book.Author, book.Quantity)
	}
	w.Flush()
	return nil
}

func (u *User) SearchBook(lib *library.Library, reader *bufio.Reader) error {
	fmt.Print("Enter the title or author of the book you want to search: ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)
	book := lib.GetBookByTitle(keyword)
	if book == nil {
		books := lib.GetBooksByAuthor(keyword)
		if len(books) > 0 {
			fmt.Println("Matching Books:")
		} else {
			return fmt.Errorf("no matching books found")
		}
		for _, book := range books {
			fmt.Printf("Title: %s, Author: %s, Quantity: %d\n", book.Title, book.Author, book.Quantity)
		}
		return nil
	} else {
		fmt.Println("Matching Books:")
		fmt.Printf("Title: %s, Author: %s, Quantity: %d\n", book.Title, book.Author, book.Quantity)
		return nil
	}

}
