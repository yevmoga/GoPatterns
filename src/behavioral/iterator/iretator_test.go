package iterator

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	shelf := new(BookShelf)

	books := []Book{
		{"Harry Potter", "J. K. Rowling"},
		{"The Lord of the Rings", "J. R. R. Tolkien"},
		{"Dunno on the Moon", "Nikolay Nosov"},
	}

	for _, book := range books {
		shelf.Add(book)
	}

	var currentBook Book
	for iterator := shelf.Iterator(); iterator.HasMore(); iterator.Next() {
		currentBook = iterator.GetCurrent()
		fmt.Println(currentBook.Author + currentBook.Name)
	}

}
