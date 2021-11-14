package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func viperConfigVariable(key string) string {
	// name of config file (without extension)
	viper.SetConfigName("config")
	// look for config in the working directory
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assersiont")
	}

	return value
}
func main() {

	name := viperConfigVariable("DB.NAME")
	fmt.Println(name)
}
