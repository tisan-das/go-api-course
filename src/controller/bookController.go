package controller

import "github.com/gin-gonic/gin"

type BookController interface {
	AddBookItem(*gin.Context)
	FetchBookItems(*gin.Context)
	FetchIndividualBookItem(*gin.Context)
	UpdateBookItem(*gin.Context)
	DeleteBookItem(*gin.Context)
}
