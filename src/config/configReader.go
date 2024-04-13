package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ConfigReader interface {
	ReadConfig() error
	GetConfigValue(string) string
	GetNestedConfigValue(string, string) string
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
