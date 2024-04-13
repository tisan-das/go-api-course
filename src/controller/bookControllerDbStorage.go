package controller

import (
	"fmt"
	"go-api-course/src/logging"
	"go-api-course/src/mapper"
	"go-api-course/src/model"
	"go-api-course/src/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookControllerDbStorage struct {
	logger   logging.Logger
	repoConn repository.Repository
}

func NewBookControllerDbStorage(logger logging.Logger, repoConn repository.Repository) BookController {
	return &BookControllerDbStorage{
		logger:   logger,
		repoConn: repoConn,
	}
}

func (svc *BookControllerDbStorage) AddBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "AddBookItem"

	var book model.Book
	err := context.BindJSON(&book)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while parsing book details: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}
	svc.logger.Debugw(fmt.Sprintf("Adding the book details %+v", book), useCase, requestId)
	storedBookDetails, err := svc.repoConn.AddBookItem(mapper.BookModelToEntityConverter(book))
	responseBook := mapper.BookEntityToModelConverter(storedBookDetails)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while storing the book details %+v: %s", book, err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(500, gin.H{"error": msg})
		return
	}
	context.JSON(201, responseBook)
}

func (svc *BookControllerDbStorage) FetchBookItems(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "FetchBookItems"

	svc.logger.Debugw("Fetching details of all the books", useCase, requestId)
	bookDetails, err := svc.repoConn.FetchAllBookItems()
	if err != nil {
		msg := fmt.Sprintf("Error occurred while fetching details of all the books: %s", err)
		svc.logger.Infow(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}
	svc.logger.Debugw(fmt.Sprintf("Fetched details of all the books: %+v", bookDetails), useCase, requestId)
	context.JSON(200, mapper.BookEntitiesToModelsConverter(bookDetails))
}

func (svc *BookControllerDbStorage) FetchIndividualBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "FetchIndividualBookItem"
	bookId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("Error occurred while parsing book id: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
	}

	svc.logger.Debugw(fmt.Sprintf("Fetching details of book with id: %d", bookId), useCase, requestId)
	bookDetails, err := svc.repoConn.FetchIndividualBookItem(bookId)
	if err != nil {
		msg := fmt.Sprintf("Error occurred fetching book details for id %d: %s", bookId, err)
		svc.logger.Infow(msg, useCase, requestId)
		context.JSON(500, gin.H{"error": msg})
		return
	}
	svc.logger.Infow(fmt.Sprintf("Fetched details of book with id %d: %+v", bookId, bookDetails), useCase, requestId)
	context.JSON(200, mapper.BookEntityToModelConverter(bookDetails))
}

func (svc *BookControllerDbStorage) UpdateBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "UpdateBookItem"
	bookId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("Error occurred while parsing book id: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}
	var book model.Book
	err = context.BindJSON(&book)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while parsing book details: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}
	if bookId != book.Id {
		msg := fmt.Sprintf("The book id of the book details is not matching with the id in path parmas")
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}

	svc.logger.Debugw(fmt.Sprintf("Updating the details of book with id %s with details %v", bookId, book), useCase, requestId)
	bookDetails, err := svc.repoConn.UpdateBookItem(bookId, mapper.BookModelToEntityConverter(book))
	if err != nil {
		msg := fmt.Sprintf("Error occurred while updating bookdetails %+v: %s", book, err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(500, gin.H{"error": msg})
		return
	}
	svc.logger.Infow(fmt.Sprintf("Updated the details of book with id %s with details %v", bookId, book), useCase, requestId)
	context.JSON(201, mapper.BookEntityToModelConverter(bookDetails))
}

func (svc *BookControllerDbStorage) DeleteBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "DeleteBookItem"
	bookId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("Error ocucrred while parsing book id: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}

	svc.logger.Debugw(fmt.Sprintf("Deleting book with id %d", bookId), useCase, requestId)
	err = svc.repoConn.DeleteBookItem(bookId)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while deleting book with id %d: %s", bookId, err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(500, gin.H{"error": msg})
		return
	}
	svc.logger.Infow(fmt.Sprintf("Deleted book with id %d", bookId), useCase, requestId)
	context.JSON(202, gin.H{"msg": fmt.Sprint("Success")})
}
