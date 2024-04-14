package model

import (
	"github.com/go-playground/validator/v10"
)

type Book struct {
	// TODO: How to discard and auto-populate ID?
	Id           int    `json:"id"`
	Name         string `json:"name" validate:"required"`
	Author       string `json:"author" validate:"required"`
	ReleasedDate string `json:"releasedDate" validate:"required"`
	ISBN         string `json:"isbn" validate:"required"`
	Price        int    `json:"price" validate:"required,gte=0"`
	Quantity     int    `json:"quantity" validate:"required,gte=0"`
}

func (book *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(book)
}

func (book *Book) Set(id int, name, author, releasedDate, isbn string, price, quantity int) {
	book.Id = id
	book.Name = name
	book.Author = author
	book.ReleasedDate = releasedDate
	book.ISBN = isbn
	book.Price = price
	book.Quantity = quantity
}
