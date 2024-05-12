package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/chegde20121/LibraryManagementSystem/internal/library"
	"github.com/chegde20121/LibraryManagementSystem/internal/user"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Welcome to the Library Management System")
		var username, password string
		fmt.Print("Enter username: ")
		fmt.Scanln(&username)
		fmt.Print("Enter password: ")
		fmt.Scanln(&password)
		curruser, err := user.Login(username, password)
		if err != nil {
			fmt.Println("Login failed:", err)
			continue
		}
		fmt.Printf("Welcome, %s!\n", curruser.UserName)
		switch curruser.Role {
		case user.ADMIN_ROLE:
			{
				for {
					fmt.Println("Select an option:")
					fmt.Println("1. Add new user")
					fmt.Println("2. Logout")
					data, _, err := reader.ReadLine()
					if err != nil {
						fmt.Println("Invalid choice. Please try again.")
						continue
					}
					choice, err := strconv.Atoi(string(data))
					if err != nil {
						fmt.Println("Invalid choice. Please try again.")
						continue
					}
					switch choice {
					case 1:
						curruser.AddNewUser(reader)
						continue
					case 2:
						curruser = nil
						fmt.Println("Logging out...")
					default:
						fmt.Println("Invalid choice. Please try again.")
						continue
					}
					if curruser == nil {
						break
					}
				}
			}
		case user.READER_ROLE:
			{
				for {
					fmt.Println("Select an option:")
					fmt.Println("1. View available books")
					fmt.Println("2. Borrow a book")
					fmt.Println("3. Return a book")
					fmt.Println("4. Search for a book")
					fmt.Println("5. View borrowed books")
					fmt.Println("6. Logout")
					var choice int
					data, _, err := reader.ReadLine()
					if err != nil {
						fmt.Println("Invalid choice. Please try again.")
						continue
					}
					choice, err = strconv.Atoi(string(data))
					if err != nil {
						fmt.Println("Invalid choice. Please try again.")
						continue
					}
					switch choice {
					case 1:
						curruser.ViewAllBooks()
					case 2:
						{

							fmt.Println("Enter the title of the book you wish to borrow:")
							data, _, err := reader.ReadLine()
							if err != nil {
								fmt.Println("Invalid input. Please try again.")
								continue
							}
							err = curruser.BorrowBook(string(data))
							if err != nil {
								fmt.Println("failed: ", err)
								continue
							}
							fmt.Println("Successfully borrowed the book", string(data))
							continue
						}
					case 3:
						{
							fmt.Print("Enter the title of the book you want to return: ")
							keyword, _ := reader.ReadString('\n')
							keyword = strings.TrimSpace(keyword)
							err := curruser.ReturnBook(keyword)
							if err != nil {
								fmt.Println("failed: ", err)
								continue
							}
							fmt.Println("Successfully returned the book")
							continue
						}
					case 4:
						{

							lib := library.GetLibraryInstance()
							err := curruser.SearchBook(lib, reader)
							if err != nil {
								fmt.Println(err)
							}
							continue
						}
					case 5:
						{
							err := curruser.DisplayBooksBorrowed()
							if err != nil {
								fmt.Println(err)
							}
							continue
						}
					case 6:
						{
							curruser = nil
							fmt.Println("Logging out...")
						}
					default:
						fmt.Println("Invalid choice. Please try again.")
					}
					if curruser == nil {
						break
					}
				}

			}
		}
	}
}
