package entity

type Book struct {
	Id           int    `gorm:"primary key;autoIncrement;not null;column:id"`
	Name         string `gorm:"not null;index:idx_name,unique;column:name"`
	Author       string `gorm:"not null;index:idx_name,unique;column:author"`
	ReleasedDate string `gorm:"index:idx_name,unique;column:releasedDate"`
	ISBN         string `gorm:"not null;index:idx_name,unique;column:isbn"`
	Price        int    `gorm:"not null;index:idx_name,unique;column:price"`
	Quantity     int    `gorm:"not null;index:idx_name,unique;column:quantity"`
}

func (book *Book) Set(id int, name, author, releasedDate, isbn string, price, quantity int) {
	// book.Id = id // Commenting this out to ensure ID would be auto populated by database
	book.Name = name
	book.Author = author
	book.ReleasedDate = releasedDate
	book.ISBN = isbn
	book.Price = price
	book.Quantity = quantity
}
