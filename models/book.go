package models

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
var Library []Book
