package entity

type Book struct {
	// TODO: How to auto populate the ID and not to accept is as part of request?
	Id           int    `gorm:"primary key;autoIncrement;column:id"`
	Name         string `gorm:"column:name"`
	Author       string `gorm:"column:author"`
	ReleasedDate string `gorm:"column:releasedDate"`
	ISBN         string `gorm:"column:isbn"`
	Price        int    `gorm:"column:price"`
	Quantity     int    `gorm:"column:quantity"`
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
