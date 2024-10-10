package main

import "os"

// a page can have a title and a body
type Page struct {
	Title string
	Body  []body
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
    filename := title + ".txt"
    body, _ := os.ReadFile(filename)
    return &Page{Title:title, Body: body}
}
