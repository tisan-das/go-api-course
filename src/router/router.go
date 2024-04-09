package router

import (
	"go-api-course/src/controller"
	"go-api-course/src/logging"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func AppRouter(logFile *os.File, bookController controller.BookController) *gin.Engine {
	router := gin.Default()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	router.Use(gin.LoggerWithFormatter(logging.FormatLogsJson))

	// TODO: Make the auth dynamic, also move the auth logic to middleware
	auth := gin.BasicAuth(gin.Accounts{
		"user1": "pass1",
		"user2": "pass2",
		"user3": "pass3",
	})

	adminGroup := router.Group("/admin", auth)
	adminGroup.POST("/book", bookController.AddBookItem)
	adminGroup.PATCH("/book/:id", bookController.UpdateBookItem)
	adminGroup.DELETE("/book/:id", bookController.DeleteBookItem)

	clientGroup := router.Group("/client")
	clientGroup.GET("/book", bookController.FetchBookItems)
	clientGroup.GET("/book/:id", bookController.FetchBookItems)

	return router
}
