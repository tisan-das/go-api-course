package controller

import (
	"fmt"
	"go-api-course/src/logging"
	"go-api-course/src/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InmemoryBookControllerService struct {
	logger logging.Logger
	books  []model.Book
}

func NewInmemoryBookController(logger logging.Logger) BookController {
	return &InmemoryBookControllerService{
		logger: logger,
		books:  make([]model.Book, 0),
	}
}

func (svc *InmemoryBookControllerService) AddBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "AddBookItem"

	var book model.Book
	err := context.BindJSON(&book)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while parsing the book details: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}
	err = book.Validate()
	if err != nil {
		msg := fmt.Sprintf("Error occurred while validating the book details %+v: %s", book, err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}
	svc.logger.Debugw(fmt.Sprintf("Processing the new book details %+v to add on the storage", book), useCase, requestId)
	var isBookExists bool = false
	for _, bookItem := range svc.books {
		if bookItem.Id == book.Id {
			isBookExists = true
			break
		}
	}
	if isBookExists {
		msg := fmt.Sprintf("A book with this id %d already exists!", book.Id)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
	} else {
		svc.books = append(svc.books, book)
		svc.logger.Infow(fmt.Sprintf("Added the new book details %+v on the storage", book), useCase, requestId)
		context.JSON(201, gin.H{"msg": fmt.Sprint("Book added")})
	}
	return
}

func (svc *InmemoryBookControllerService) FetchBookItems(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "FetchBookItems"
	// TODO: How to introduce pagination?
	svc.logger.Debugw("Fetching all the books", useCase, requestId)
	books := svc.books
	svc.logger.Infow(fmt.Sprintf("Fetched list of books to return: %+v", books), useCase, requestId)
	context.JSON(200, books)
}

func (svc *InmemoryBookControllerService) FetchIndividualBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "FetchIndividualBookItem"

	bookId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("Error occurred while fetching book id from path params: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(500, gin.H{"error": msg})
		return
	}
	svc.logger.Debugw(fmt.Sprintf("Fetching book details for the book with id: %d", bookId), useCase, requestId)
	var book model.Book
	for _, bookItem := range svc.books {
		if bookItem.Id == bookId {
			book = bookItem
			break
		}
	}
	svc.logger.Infow(fmt.Sprintf("Fetched details for book with id %d: %+v", bookId, book), useCase, requestId)
	if book.Id != 0 {
		context.JSON(200, book)
	} else {
		context.JSON(404, gin.H{"error": fmt.Sprintf("Book with id %d is not found", bookId)})
	}
	return
}

func (svc *InmemoryBookControllerService) UpdateBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "UpdateBookItem"

	bookId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("Error occurred while fetching book id from path params: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(500, gin.H{"error": msg})
		return
	}
	var book model.Book
	err = context.BindJSON(&book)
	if err != nil {
		msg := fmt.Sprintf("Error occurred while parsing the book details: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}
	err = book.Validate()
	if err != nil {
		msg := fmt.Sprintf("Error occurred while validating the book details %+v: %s", book, err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(400, gin.H{"error": msg})
		return
	}
	svc.logger.Debugw(fmt.Sprintf("Updating the new book details %+v with id %d to on the storage", book, bookId), useCase, requestId)

	var bookIndex int = -1
	for index, bookItem := range svc.books {
		if bookItem.Id == bookId {
			bookIndex = index
			break
		}
	}
	if bookIndex == -1 {
		context.JSON(404, gin.H{"error": fmt.Sprintf("Book with id %d is not found", bookId)})
	} else {
		svc.books[bookIndex] = book
		context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
	}
	return
}

func (svc *InmemoryBookControllerService) DeleteBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "DeleteBookItem"

	bookId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		msg := fmt.Sprintf("Error occurred while fetching book id from path params: %s", err)
		svc.logger.Errorw(msg, useCase, requestId)
		context.JSON(500, gin.H{"error": msg})
		return
	}
	svc.logger.Debugw(fmt.Sprintf("Removing book details for the book with id: %d", bookId), useCase, requestId)
	var bookIndex int = -1
	for index, bookItem := range svc.books {
		if bookItem.Id == bookId {
			bookIndex = index
			break
		}
	}
	svc.logger.Infow(fmt.Sprintf("Fetched index for book with id %d: %d", bookId, bookIndex), useCase, requestId)
	if bookIndex == -1 {
		context.JSON(404, gin.H{"error": fmt.Sprintf("Book with id %d is not found", bookId)})
	} else {
		svc.books = append(svc.books[:bookIndex], svc.books[bookIndex+1:]...)
		context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
	}
	return
}
