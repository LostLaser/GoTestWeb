package models

import "os"

// Book contains information relevant to an individual book
type Book struct {
	Title string
	ISBN  string
	Genre string
}

// BookPageData is a response body for books
type BookPageData struct {
	PageTitle string
	Library   []Book
}

// Library contains all of the books
var Library map[string]Book

// DeleteBook removes the first item found with specified ISBN
func DeleteBook(isbn string) error {

	if _, ok := Library[isbn]; ok {
		delete(Library, isbn)
	} else {
		return os.ErrNotExist
	}

	return nil
}

// AddBook adds the input book to the library
func AddBook(inputBook Book) error {
	if inputBook.ISBN == "" {
		return os.ErrInvalid
	}
	if _, ok := Library[inputBook.ISBN]; !ok {
		return os.ErrExist
	}

	Library[inputBook.ISBN] = inputBook

	return nil
}

// GetBook retrieves a single book from the library
func GetBook(isbn string) (Book, error) {
	if val, ok := Library[isbn]; ok {
		return val, nil
	}

	return Book{}, os.ErrNotExist
}

// GetBooks retrieves the entire library
func GetBooks() map[string]Book {
	return Library
}
