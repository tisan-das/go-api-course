package model

import (
	"github.com/go-playground/validator/v10"
)

type Book struct {
	Id           int    `json:"Id" validate:"required,gt=0"`
	Name         string `json:"Name" validate:"required"`
	Author       string `json:"Author" validate:"required"`
	ReleasedDate string `json:"ReleasedDate" validate:"required"`
	ISBN         string `json:"ISBN" validate:"required"`
	Price        int    `json:"Price" validate:"required,gte=0"`
	Quantity     int    `json:"Quantity" validate:"required,gte=0"`
}

func (book *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(book)
}
