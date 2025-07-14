package controllers

import (
	"bufio"
	"fmt"
	"libaray_management/models"
	"libaray_management/services"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	manager services.LibraryManger
}

func New() *LibraryController {
	return &LibraryController{
		manager: services.New(),
	}
}

func (l *LibraryController) AddMember() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	memberName, _ := reader.ReadString('\n')
	memberName = strings.TrimSpace(memberName)

	member := models.Member{
		ID:       rand.Int() % 50,
		Name:     memberName,
		Borrowed: []models.Book{},
	}
	l.manager.AddMember(member)
	fmt.Printf("Member have beed added with id = %v\n", member.ID)
}

func (l *LibraryController) AddBook() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	bookCnt := l.manager.BookCount()
	book := models.Book{
		ID:     bookCnt,
		Title:  title,
		Author: author,
		Status: models.BookStatusAvailable,
	}

	l.manager.AddBook(book)
	fmt.Printf("Book was added successfully. Its ID = %v\n", book.ID)
}

func (l *LibraryController) RemoveBook() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	bookId, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID")
		return
	}

	l.manager.RemoveBook(bookId)
	fmt.Println("Removed the book with ID =", bookId)
}

func (l *LibraryController) BorrowBook() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter book ID to borrow: ")
	bookStr, _ := reader.ReadString('\n')
	bookId, err := strconv.Atoi(strings.TrimSpace(bookStr))
	if err != nil {
		fmt.Println("Invalid book ID")
		return
	}

	fmt.Print("Enter your member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberId, err := strconv.Atoi(strings.TrimSpace(memberStr))
	if err != nil {
		fmt.Println("Invalid member ID")
		return
	}

	if err := l.manager.BorrowBook(bookId, memberId); err != nil {
		fmt.Printf("Unable to borrow book: %v\n", err.Error())
		return
	}

	fmt.Println("Successfully borrowed the book")
}

func (l *LibraryController) ReturnBook() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter book ID to return: ")
	bookStr, _ := reader.ReadString('\n')
	bookId, err := strconv.Atoi(strings.TrimSpace(bookStr))
	if err != nil {
		fmt.Println("Invalid book ID")
		return
	}

	fmt.Print("Enter your member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberId, err := strconv.Atoi(strings.TrimSpace(memberStr))
	if err != nil {
		fmt.Println("Invalid member ID")
		return
	}

	if err := l.manager.ReturnBook(bookId, memberId); err != nil {
		fmt.Printf("Unable to return book: %v\n", err.Error())
		return
	}

	fmt.Println("Successfully returned the book")
}

func (l *LibraryController) ListAvailableBooks() {
	books := l.manager.ListAvailableBooks()

	if len(books) == 0 {
		fmt.Println("There are no available books")
		return
	}

	fmt.Println("Available books:")
	for idx, book := range books {
		fmt.Printf("%v. Title: %v, Author: %v, ID: %v\n", idx+1, book.Title, book.Author, book.ID)
	}
	fmt.Println("--------------------------------------------------")
}

func (l *LibraryController) ListBorrowedBooks() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your member ID: ")
	memberStr, _ := reader.ReadString('\n')
	memberId, err := strconv.Atoi(strings.TrimSpace(memberStr))
	if err != nil {
		fmt.Println("Invalid member ID")
		return
	}

	books := l.manager.ListBorrowedBooks(memberId)

	if len(books) == 0 {
		fmt.Println("You have not borrowed any books")
		return
	}

	fmt.Println("Borrowed books:")
	for idx, book := range books {
		fmt.Printf("%v. Title: %v, Author: %v, ID: %v\n", idx+1, book.Title, book.Author, book.ID)
	}
	fmt.Println("--------------------------------------------------")
}
