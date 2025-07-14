package models

type BookStatus string

const BookStatusAvailable = BookStatus("Available")
const BookStatusBorrowed = BookStatus("Borrowed")

type Book struct {
	ID     int
	Title  string
	Author string
	Status BookStatus
}
