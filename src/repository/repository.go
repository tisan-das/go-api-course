package repository

import "go-api-course/src/entity"

type Repository interface {
	InitDBConnection(dbHost, dbName, dbUser, dbPassword, dbPort string) error
	AutoMigration() error
	AddBookItem(entity.Book) (entity.Book, error)
	FetchIndividualBookItem(id int) (entity.Book, error)
	FetchAllBookItems() ([]entity.Book, error)
	UpdateBookItem(int, entity.Book) (entity.Book, error)
	DeleteBookItem(int) error
}
