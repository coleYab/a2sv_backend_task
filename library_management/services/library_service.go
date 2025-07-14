package services

import (
	"fmt"
	"libaray_management/models"
)

type LibraryManger interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
	BookCount() int
	AddMember(member models.Member)
}

type LibraryService struct {
	books     map[int]*models.Book
	bookCount int
	members   map[int]*models.Member
}

func New() *LibraryService {
	return &LibraryService{
		bookCount: 0,
		books:     map[int]*models.Book{},
		members:   map[int]*models.Member{},
	}
}

func (l *LibraryService) AddMember(member models.Member) {
	l.members[member.ID] = &member
}

func (l *LibraryService) ReturnBook(bookID int, memeberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return fmt.Errorf("book not found")
	}

	member, ok := l.members[memeberID]
	if !ok {
		return fmt.Errorf("member not found")
	}

	for idx, borrowed := range member.Borrowed {
		if borrowed.ID == bookID {
			after := append(member.Borrowed[:idx], member.Borrowed[idx+1:]...)
			member.Borrowed = after
			book.Status = models.BookStatusAvailable
			return nil
		}
	}

	return fmt.Errorf("book was not borrowed by you")
}

func (l *LibraryService) BorrowBook(bookID int, memberID int) error {
	book, ok := l.books[bookID]
	if !ok {
		return fmt.Errorf("book not found")
	}

	member, ok := l.members[memberID]
	if !ok {
		return fmt.Errorf("member is not found")
	}

	if book.Status == models.BookStatusBorrowed {
		return fmt.Errorf("book is already borrowed")
	}

	member.Borrowed = append(member.Borrowed, *book)
	book.Status = models.BookStatusBorrowed
	return nil
}

func (l *LibraryService) BookCount() int {
	return l.bookCount
}

func (l *LibraryService) AddBook(book models.Book) {
	l.books[book.ID] = &book
	l.bookCount += 1
}

func (l *LibraryService) RemoveBook(bookId int) {
	book, ok := l.books[bookId]
	if !ok {
		return
	}

	if book.Status == models.BookStatusBorrowed {
		// delete it from the borrowed book
		for _, member := range l.members {
			found := false
			for idx, book := range member.Borrowed {
				if book.ID == bookId {
					final := append(member.Borrowed[:idx], member.Borrowed[idx+1:]...)
					member.Borrowed = final
					found = true
					break
				}
			}
			if found {
				break
			}
		}
	}

	delete(l.books, bookId)
}

func (l LibraryService) ListAvailableBooks() []models.Book {
	books := []models.Book{}

	for _, book := range l.books {
		if book.Status == models.BookStatusAvailable {
			books = append(books, *book)
		}
	}

	return books
}

func (l LibraryService) ListBorrowedBooks(memeberID int) []models.Book {
	member, ok := l.members[memeberID]
	if !ok {
		return []models.Book{}
	}

	return member.Borrowed
}
