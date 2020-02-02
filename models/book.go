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
	Library   map[string]Book
}

var library = make(map[string]Book)

// DeleteBook removes the first item found with specified ISBN
func DeleteBook(isbn string) error {

	if _, ok := library[isbn]; ok {
		delete(library, isbn)
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
	if _, found := library[inputBook.ISBN]; found {
		return os.ErrExist
	}

	library[inputBook.ISBN] = inputBook

	return nil
}

// GetBook retrieves a single book from the library
func GetBook(isbn string) (Book, error) {
	if val, found := library[isbn]; !found {
		return val, nil
	}

	return Book{}, os.ErrNotExist
}

// GetBooks retrieves the entire library
func GetBooks() map[string]Book {
	return library
}
