package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func getEnvViperVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type asseration")
	}
	return value
}

func main() {
	viperValue := getEnvViperVariable("BASEURL")
	fmt.Println("Hello World")
	fmt.Println(viperValue)
}
