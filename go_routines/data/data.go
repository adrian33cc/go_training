package data

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "Dune", false},
	{2, "Harry Potter 1", false},
	{3, "Harry Potter 2", false},
	{4, "Harry Potter 3", false},
	{5, "Harry Potter 4", false},
	{6, "Harry Potter 5", false},
	{7, "Harry Potter 6", false},
	{8, "Harry Potter 7", false},
	{9, "Piratas del Caribe 1", false},
	{10, "Piratas del Caribe 2", false},
}

func findBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1

	var book *Book

	m.RLock()
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}

	m.RUnlock()

	return index, book

}

func FinishedBook(id int, m *sync.RWMutex) {
	i, book := findBook(id, m)

	if i < 0 {
		return
	}

	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()
	fmt.Printf("Finished Book: %s \n", book.Title)
}
