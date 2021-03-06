package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintln(b.Title)
}

var books = []Book{
	{
		ID:            1,
		Title:         "Book1",
		Author:        "me",
		YearPublished: 1990,
	},
	{
		ID:            2,
		Title:         "Book2",
		Author:        "you",
		YearPublished: 1991,
	},
}
