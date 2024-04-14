package config

import (
	"fmt"
	"go-api-course/src/logging"
	"go-api-course/src/repository"
	"strconv"
)

type AppConfig struct {
	logger       logging.Logger
	repoConn     repository.Repository
	configReader ConfigReader
}

func NewAppConfig(logger logging.Logger, repoConn repository.Repository) AppConfig {

	var appConfig AppConfig = AppConfig{
		logger:       logger,
		repoConn:     repoConn,
		configReader: NewViperConfigReader(),
	}
	return appConfig
}

func (appConfig *AppConfig) InitializeApplicationConfiguration() error {

	// Read the configuration
	err := appConfig.configReader.ReadConfig()
	if err != nil {
		fmt.Println("Error occurred while reading configuration: ", err)
		panic(err)
	}

	loggingLevel := appConfig.configReader.GetConfigValue("logLevel")
	port, err := strconv.Atoi(appConfig.configReader.GetConfigValue("port"))
	if err != nil {
		msg := fmt.Sprintf("Error occurred while capturing port value: %s", err)
		return fmt.Errorf("%s", msg)
	}
	logFileName := appConfig.configReader.GetConfigValue("logFile")
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
		msg := fmt.Sprintf("Error occurred while initializing application logger: %s", err)
		return fmt.Errorf("%s", msg)
	}
	appConfig.logger = logger

	// Initialize DB connection
	var repositoryConnection repository.Repository
	dbUser := appConfig.configReader.GetNestedConfigValue("database", "user")
	dbPassword := appConfig.configReader.GetNestedConfigValue("database", "password")
	dbName := appConfig.configReader.GetNestedConfigValue("database", "name")
	dbPort := appConfig.configReader.GetNestedConfigValue("database", "port")
	dbHost := appConfig.configReader.GetNestedConfigValue("database", "host")
	if appConfig.repoConn == nil {
		repositoryConnection, err = repository.NewPostgresRepo(dbHost, dbName, dbUser,
			dbPassword, dbPort)
		appConfig.repoConn = repositoryConnection

		if err != nil {
			msg := fmt.Sprintf("Error occurred while initiating DB connection: %s", err)
			return fmt.Errorf("%s", msg)
		}
		err = repositoryConnection.AutoMigration()
		if err != nil {
			msg := fmt.Sprintf("Error occurred migrating repo: %s", err)
			return fmt.Errorf("%s", msg)
		}
	} else {
		err = appConfig.repoConn.InitDBConnection(dbHost, dbName, dbUser, dbPassword, dbPort)
		if err != nil {
			msg := fmt.Sprintf("Error occurred while initiating DB connection: %s", err)
			return fmt.Errorf("%s", msg)
		}
	}

	return nil
}

func (appConfig *AppConfig) WatchConfigChange() {
	appConfig.configReader.SetWatchConfig(appConfig)
}

func (appConfig *AppConfig) GetConfigValue(str string) string {
	return appConfig.configReader.GetConfigValue(str)
}

func (appConfig *AppConfig) GetLogger() logging.Logger {
	return appConfig.logger
}

func (appConfig *AppConfig) GetRepositoryConnection() repository.Repository {
	return appConfig.repoConn
}
