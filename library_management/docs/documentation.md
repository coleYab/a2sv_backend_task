# 📚 Library Management System (Go CLI)
This is a simple command-line Library Management System implemented in Go. It allows users to manage a small-scale library, including adding/removing books, borrowing/returning books, and listing available or borrowed books. The system also supports member registration at the beginning of the session.

## 🚀 How to Run
Ensure Go is installed: go version

Clone the repository or place the code inside a Go module.

Run the application:
```bash
cd library_management
go run main.go
```

## 📜 Program Flow
### 1. Member Registration
Upon launch, the application prompts to register a library member via: `controller.AddMember()`

### 2. Menu Options
The user is presented with an interactive CLI menu to choose an action:
```bash
go run main.go
1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books
0. Exit
The menu runs in an infinite loop until the user chooses to exit (0).
```

### 🔧 Features
`controller.AddBook()` - Adds a new book to the library system.

`controller.RemoveBook()` - Removes a specified book from the system.

`controller.BorrowBook()` - Allows the registered member to borrow a book.

`controller.ReturnBook()` - Enables a member to return a previously borrowed book.

`controller.ListAvailableBooks()` - Displays all books currently available for borrowing.

`controller.ListBorrowedBooks()` - Displays all books currently borrowed by the member(s).

### 🔤 Example Output
```bash
go run main.go
===   Welcome to library mangement system ====
Enter member name: Alice

Library Menu:
1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books
0. Exit
Enter your choice: 1
Enter book title: The Go Programming Language
Book added successfully.
```
