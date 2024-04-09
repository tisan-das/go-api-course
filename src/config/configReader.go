package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ConfigReader interface {
	ReadConfig() error
	GetConfigValue(string) string
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
