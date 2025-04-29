package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Collection interface {
	CreateIterator() Iterator
}

type BookCollection struct {
	titles []string
}

type BookIterator struct {
	books []string
	index int
}

func (bi *BookIterator) HasNext() bool {
	return bi.index < len(bi.books)
}

func (bi *BookIterator) Next() interface{} {
	if bi.HasNext() {
		book := bi.books[bi.index]
		bi.index++
		return book
	}
	return nil
}

func (bc BookCollection) CreateIterator() Iterator {
	return &BookIterator{
		books: bc.titles,
		index: 0,
	}
}

func main() {
	library := &BookCollection{titles: []string{"Clean Code", "Design Patterns", "Golang in Action"}}
	iterator := library.CreateIterator()
	for iterator.HasNext() {
		book := iterator.Next().(string)
		fmt.Println("Book: ", book)

	}
}
