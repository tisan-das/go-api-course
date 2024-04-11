package controller

import (
	"fmt"
	"go-api-course/src/logging"

	"github.com/gin-gonic/gin"
)

type InmemoryBookControllerService struct {
	logger logging.Logger
}

func NewInmemoryBookController(logger logging.Logger) BookController {
	return &InmemoryBookControllerService{
		logger: logger,
	}
}

func (svc *InmemoryBookControllerService) AddBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "AddBookItem"
	svc.logger.Infow("---> Implementation pending for adding book item!", useCase, requestId)
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}

func (svc *InmemoryBookControllerService) FetchBookItems(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "FetchBookItems"
	svc.logger.Infow("---> Implementation pending for fetching all the book items!", useCase, requestId)
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}

func (svc *InmemoryBookControllerService) FetchIndividualBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "FetchIndividualBookItem"
	svc.logger.Infow("---> Implementation pending for fetching individual book item!", useCase, requestId)
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}

func (svc *InmemoryBookControllerService) UpdateBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "UpdateBookItem"
	svc.logger.Infow("---> Implementation pending for updating book item!", useCase, requestId)
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}

func (svc *InmemoryBookControllerService) DeleteBookItem(context *gin.Context) {
	requestIdValue, _ := context.Get("requestId")
	requestId := requestIdValue.(string)
	useCase := "DeleteBookItem"
	svc.logger.Infow("---> Implementation pending for deleting book item!", useCase, requestId)
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}
