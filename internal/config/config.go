package config

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

var runtimeViper = viper.New()

func init() {
	fmt.Println("initial Config")
}

type Config struct {
	Port     int
	Database struct {
		Type string
		URL  string
	}
	Loggin struct {
		Path string
	}
}

func (d Config) String() string {
	b, err := json.Marshal(d)
	if err != nil {
		panic(err.Error())
	}
	return string(b)
}

func allConf() (*Config, error) {
	c := &Config{}
	err := runtimeViper.Unmarshal(c)
	return c, err
}

func Viper() *viper.Viper {
	return runtimeViper
}

func Get(key string) string {
	val := runtimeViper.Get(key)
	switch val.(type) {
	case string:
		return val.(string)
	case int:
		return strconv.Itoa(val.(int))
	case float64:
		return fmt.Sprintf("%f", val.(float64))
	default:
		return ""
	}
}
