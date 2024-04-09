package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type InmemoryBookControllerService struct{}

func NewInmemoryBookController() BookController {
	return &InmemoryBookControllerService{}
}

func (svc *InmemoryBookControllerService) AddBookItem(context *gin.Context) {
	fmt.Println("---> Implementation pending for adding book item!")
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}

func (svc *InmemoryBookControllerService) FetchBookItems(context *gin.Context) {
	fmt.Println("---> Implementation pending for fetching all the book items!")
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}

func (svc *InmemoryBookControllerService) FetchIndividualBookItem(context *gin.Context) {
	fmt.Println("---> Implementation pending for fetching individual book item!")
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}

func (svc *InmemoryBookControllerService) UpdateBookItem(context *gin.Context) {
	fmt.Println("---> Implementation pending for updating book item!")
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}

func (svc *InmemoryBookControllerService) DeleteBookItem(context *gin.Context) {
	fmt.Println("---> Implementation pending for deleting book item!")
	context.JSON(200, gin.H{"msg": fmt.Sprint("Success")})
}
