package main

import (
	"fmt"
	"go-api-course/src/config"
	"go-api-course/src/controller"
	"go-api-course/src/router"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize Application Configurations
	os.Setenv("config", "dev")
	appConfig := config.NewAppConfig(nil, nil)
	err := appConfig.InitializeApplicationConfiguration()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	appConfig.WatchConfigChange()
	logger := appConfig.GetLogger()
	repositoryConnection := appConfig.GetRepositoryConnection()
	serverPort := appConfig.GetConfigValue("port")

	// Initialize Router and Controllers
	var bookController controller.BookController
	// bookController = controller.NewInmemoryBookController(logger)
	bookController = controller.NewBookControllerDbStorage(logger, repositoryConnection)

	router := router.AppRouter(bookController, logger)
	// router.Run(":4040")
	server := http.Server{
		Addr:         fmt.Sprint(":", serverPort),
		Handler:      router,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// TODO: Add cron scheduler for example
	// TODO: Add some workflow for example
	server.ListenAndServe()
}

func getData(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{"msg": "hello!", "name": name})
}

func postData(c *gin.Context) {
	body := c.Request.Body
	value, _ := io.ReadAll(body)
	c.JSON(200, gin.H{"msg": fmt.Sprint("hello ", string(value), "!")})
}
