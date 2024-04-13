package main

import (
	"fmt"
	"go-api-course/src/config"
	"go-api-course/src/controller"
	"go-api-course/src/logging"
	"go-api-course/src/repository"
	"go-api-course/src/router"
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
		panic(err)
	}

	loggingLevel := configReader.GetConfigValue("logLevel")
	port, err := strconv.Atoi(configReader.GetConfigValue("port"))
	if err != nil {
		fmt.Println("Error occurred while capturing port value: ", err)
		panic(err)
	}
	logFileName := configReader.GetConfigValue("logFile")
	fmt.Println(loggingLevel, port, logFileName)

	// Initialize App loggers
	var logger logging.Logger
	if loggingLevel == string(logging.LOG_INFO_LEVEL) {
		logger, err = logging.NewZapSugarLogger(logging.LOG_INFO_LEVEL, logFileName)
	} else if loggingLevel == string(logging.LOG_WARN_LEVEL) {
		logger, err = logging.NewZapSugarLogger(logging.LOG_WARN_LEVEL, logFileName)
	} else {
		logger, err = logging.NewZapSugarLogger(logging.LOG_DEBUG_LEVEL, logFileName)
	}
	if logger == nil || err != nil {
		fmt.Println("Error occurred while initializing application logger: ", err)
		panic(err)
	}

	// Initialize DB connection
	var repositoryConnection repository.Repository
	dbUser := configReader.GetNestedConfigValue("database", "user")
	dbPassword := configReader.GetNestedConfigValue("database", "password")
	dbName := configReader.GetNestedConfigValue("database", "name")
	dbPort := configReader.GetNestedConfigValue("database", "port")
	dbHost := configReader.GetNestedConfigValue("database", "host")
	repositoryConnection, err = repository.NewPostgresRepo(dbHost, dbName, dbUser,
		dbPassword, dbPort)
	if err != nil {
		fmt.Println("Error occurred while initiating DB connection: ", err)
		panic(err)
	}
	err = repositoryConnection.AutoMigration()
	if err != nil {
		fmt.Println("Error occurred migrating repo: ", err)
		panic(err)
	}

	// Initialize Router and Controllers
	var bookController controller.BookController
	// bookController = controller.NewInmemoryBookController(logger)
	bookController = controller.NewBookControllerDbStorage(logger, repositoryConnection)

	router := router.AppRouter(bookController, logger)
	// router.Run(":4040")
	server := http.Server{
		Addr:         fmt.Sprint(":", port),
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
