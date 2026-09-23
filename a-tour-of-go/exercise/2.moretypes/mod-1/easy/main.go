package main

import "fmt"

type Book struct {
	Title string
	Pages int
	Read  bool
}

var books = [3]Book{
	{Title: "Go Basics", Pages: 100, Read: false},
	{Title: "Short Stories", Pages: 80, Read: false},
	{Title: "Field Notes", Pages: 40, Read: true},
}

var selected = books[:2]

func markRead(book *Book) {
	book.Read = true
}

func addPages(book *Book, extra int) bool {
	if extra <= 0 {
		return false
	}

	book.Pages += extra
	return true
}

func main() {
	markRead(&selected[0])
	addPages(&selected[0], 20)
	fmt.Println(books[:1], books[0])
	fmt.Printf("len=%d cap=%d \n", len(selected), cap(selected))
}

//Why does a change through a pointer to `selected[0]`
// also change the first array element?
//
// Because selected is pointed to the Books address instead of just a copy.
// So changing the selected will also change the original data.
