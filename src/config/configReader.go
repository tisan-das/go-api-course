package config

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigReader interface {
	ReadConfig() error
	GetConfigValue(string) string
	GetNestedConfigValue(string, string) string
	SetWatchConfig(*AppConfig)
}

type ViperConfigReader struct{}

func NewViperConfigReader() ConfigReader {
	return &ViperConfigReader{}
}

func (configReader *ViperConfigReader) ReadConfig() error {
	configFileName := fmt.Sprint(os.Getenv("config"), ".json")
	fmt.Println("Reading the configuration from file :", configFileName)
	viper.SetConfigName(os.Getenv("config"))
	viper.SetConfigType("json")
	viper.AddConfigPath("src/config/")
	viper.AddConfigPath(".")
	return viper.ReadInConfig()
}

func (configReader *ViperConfigReader) GetConfigValue(key string) string {
	fmt.Println("Reading key: ", key)
	return viper.Get(key).(string)
}

func (configReader *ViperConfigReader) GetNestedConfigValue(key1, key2 string) string {
	fmt.Println("Reading key: (", key1, key2, ")")
	var value map[string]interface{}
	value = viper.Get(key1).(map[string]interface{})
	return value[key2].(string)
}

func (configReader *ViperConfigReader) SetWatchConfig(appConfig *AppConfig) {
	var err error
	viper.WatchConfig()
	viper.OnConfigChange(func(event fsnotify.Event) {
		appConfig.logger.Infow(fmt.Sprintf("Config change detected: %s", event.Name), "configChange", "")
		err = appConfig.InitializeApplicationConfiguration()
		if err != nil {
			msg := fmt.Sprintf("Error occurred while reinitializing config: %s", err)
			appConfig.logger.Errorw(msg, "configChange", "")
			// Triggering panic for application would let us know if there's some issue with DB connection
			panic(msg)
		}
	})
}
