package main

import (
	"fmt"
	"library_management/controllers"
)

func main() {
	controller := controllers.New()

	fmt.Println("===   Welcome to library mangement system ====")
	controller.AddMember()

	for {
		fmt.Println("\nLibrary Menu:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("0. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			controller.AddBook()
		case 2:
			controller.RemoveBook()
		case 3:
			controller.BorrowBook()
		case 4:
			controller.ReturnBook()
		case 5:
			controller.ListAvailableBooks()
		case 6:
			controller.ListBorrowedBooks()
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
