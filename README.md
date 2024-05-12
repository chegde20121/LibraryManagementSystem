## Introduction

This is a console-based Library Management System developed in Go. The system allows users to manage the inventory of books in a library, including borrowing, returning, searching, and displaying available books.

## Features

- View the list of available books.
- Borrow a book from the library.
- Return a book to the library.
- Search for a book by title or author name.
- Display borrowed books.

## Usage:-

### Admin Role Actions

Admin users have access to the following actions:

#### 1. Add New User

Admin can add a new user to the system by selecting the "Add new user" option. They will be prompted to enter the details of the new user, including username, password, and role.

###### 2. Logout

Admin can choose to logout from the system by selecting the "Logout" option. This will log out the current admin user and return to the login screen.

#### Usage Example:


- **Add new user:** Admin selects option 1 and follows the prompts to add a new user.
- **Logout:** Admin selects option 2 to logout from the system.

#### Additional Information

- If the admin enters an invalid choice, they will be prompted to try again until a valid choice is made.
- After performing an action, the admin will be returned to the main menu to select another option or logout.
- The admin's actions are restricted to the options provided, ensuring secure and controlled access to the system.

### Reader Role Actions

Users with the reader role have access to the following actions:

#### 1. View Available Books

Readers can view all the available books in the library. This includes information such as the title, author, and quantity of each book.

#### 2. Borrow a Book

Readers can borrow a book from the library by selecting the "Borrow a book" option. They will be prompted to enter the title of the book they wish to borrow.

#### 3. Return a Book

Readers can return a borrowed book to the library by selecting the "Return a book" option. They will be prompted to enter the title of the book they want to return.

#### 4. Search for a Book

Readers can search for books by title or author name using the "Search for a book" option. They will be prompted to enter the title or author of the book they want to search for.

#### 5. View Borrowed Books

Readers can view a list of books they have currently borrowed from the library using the "View borrowed books" option.

#### 6. Logout

Readers can choose to logout from the system by selecting the "Logout" option. This will log out the current reader user and return to the login screen.

#### Usage Example:


- **View available books:** Reader selects option 1 to view all available books.
- **Borrow a book:** Reader selects option 2 and enters the title of the book they wish to borrow.
- **Return a book:** Reader selects option 3 and enters the title of the book they want to return.
- **Search for a book:** Reader selects option 4 and enters the title or author of the book they want to search for.
- **View borrowed books:** Reader selects option 5 to view a list of books they have borrowed.
- **Logout:** Reader selects option 6 to logout from the system.

#### Additional Information

- If the reader enters an invalid choice, they will be prompted to try again until a valid choice is made.
- After performing an action, the reader will be returned to the main menu to select another option or logout.
- The reader's actions are restricted to the options provided, ensuring secure and controlled access to the system.

### Prerequisites

- Go programming language installed on your system.

### Run Application

- you can use the shell script "build.sh" for running script
- Before running the script you might have to  give the necessary permission
```
chmod +x build.sh

./build.sh
```

### Features

- Build the project with unit tests and display coverage.
- Build and run the application.
- Build the application without running unit tests.