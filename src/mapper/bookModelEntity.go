package mapper

import (
	"go-api-course/src/entity"
	"go-api-course/src/model"
)

func BookEntityToModelConverter(book entity.Book) model.Book {
	var bookModel model.Book
	bookModel.Set(book.Id, book.Name, book.Author, book.ReleasedDate, book.ISBN,
		book.Price, book.Quantity)
	return bookModel
}

func BookEntitiesToModelsConverter(books []entity.Book) []model.Book {
	var bookModels []model.Book
	bookModels = make([]model.Book, 0)
	for _, bookItem := range books {
		var bookModel model.Book
		bookModel.Set(bookItem.Id, bookItem.Name, bookItem.Author,
			bookItem.ReleasedDate, bookItem.ISBN, bookItem.Price, bookItem.Quantity)
		bookModels = append(bookModels, bookModel)
	}
	return bookModels
}

func BookModelToEntityConverter(book model.Book) entity.Book {
	var bookEntity entity.Book
	bookEntity.Set(book.Id, book.Name, book.Author, book.ReleasedDate, book.ISBN,
		book.Price, book.Quantity)
	return bookEntity
}
