package repository

import (
	"fmt"
	"go-api-course/src/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
)

type PostgresRepo struct {
	DB *gorm.DB
}

func NewPostgresRepo(dbHost, dbName, dbUser, dbPassword, dbPort string) (Repository,
	error) {
	var postgresRepo PostgresRepo
	err := postgresRepo.InitDBConnection(dbHost, dbName, dbUser, dbPassword, dbPort)
	return &postgresRepo, err
}

// TODO: How to auto-rotate password?
// TODO: How to omit the logs for credential issues?
// TODO: How to use singleton design pattern for DB connection?
func (repo *PostgresRepo) InitDBConnection(dbHost, dbName, dbUser, dbPassword,
	dbPort string) error {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost,
		dbUser, dbPassword, dbName, dbPort)
	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "book_service.",
			SingularTable: false,
		},
	})
	repo.DB = dbConn
	return err
}

func (repo *PostgresRepo) AutoMigration() error {
	return repo.DB.AutoMigrate(&entity.Book{})
}

func (repo *PostgresRepo) AddBookItem(book entity.Book) (entity.Book, error) {
	result := repo.DB.Clauses(clause.Returning{}).Create(&book)
	if result.Error != nil {
		return book, result.Error
	}
	if result.RowsAffected == 0 {
		msg := "No rows created"
		return book, fmt.Errorf("%s", msg)
	}
	return book, nil
}

func (repo *PostgresRepo) FetchIndividualBookItem(id int) (entity.Book, error) {
	var book entity.Book
	result := repo.DB.First(&book, id)
	if result.Error != nil {
		return book, result.Error
	}
	return book, nil
}

func (repo *PostgresRepo) FetchAllBookItems() ([]entity.Book, error) {
	var books []entity.Book
	result := repo.DB.Find(&books)
	if result.Error != nil {
		return books, result.Error
	}
	return books, nil
}

func (repo *PostgresRepo) UpdateBookItem(id int, book entity.Book) (entity.Book, error) {
	book.Id = id
	result := repo.DB.Model(entity.Book{}).Where("id", id).Updates(book)
	if result.Error != nil {
		return book, result.Error
	}
	if result.RowsAffected == 0 {
		msg := "No rows updated"
		return book, fmt.Errorf("%s", msg)
	}
	return book, nil
}

func (repo *PostgresRepo) DeleteBookItem(id int) error {
	result := repo.DB.Delete(entity.Book{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		msg := "No rows deleted"
		return fmt.Errorf("%s", msg)
	}
	return nil
}
