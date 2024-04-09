package main

import (
	"fmt"
	"go-api-course/src/config"
	"go-api-course/src/logging"
	"go-api-course/src/middleware"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize App configurations
	var configReader config.ConfigReader
	os.Setenv("config", "dev")
	configReader = config.NewViperConfigReader()
	err := configReader.ReadConfig()
	if err != nil {
		fmt.Println("Error occurred while reading configuration: ", err)
		return
	}

	loggingLevel := configReader.GetConfigValue("loggingLevel")
	port, err := strconv.Atoi(configReader.GetConfigValue("port"))
	if err != nil {
		fmt.Println("Error occurred while capturing port value: ", err)
		return
	}
	fmt.Println(loggingLevel, port)

	// gin.BasicAuth
	fmt.Println("hello!")
	router := gin.Default()

	loggingFile, err := os.Create("output.log")
	if err != nil {
		fmt.Println("Error occurred initiating log file: ", err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(loggingFile, os.Stdout)
	router.Use(gin.LoggerWithFormatter(logging.FormatLogsJson))

	auth := gin.BasicAuth(gin.Accounts{
		"user1": "pass1",
		"user2": "pass2",
		"user3": "pass3",
	})

	fmt.Println(auth)
	// adminGroup := router.Group("/admin", middleware.Authenticate,auth)
	adminGroup := router.Group("/admin", middleware.Authenticate, middleware.AddHeader)
	adminGroup.GET("/data", getData)
	clientGroup := router.Group("/client")
	clientGroup.POST("/data", postData)

	// router.Run(":4040")
	server := http.Server{
		Addr:         fmt.Sprint(":", port),
		Handler:      router,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
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
