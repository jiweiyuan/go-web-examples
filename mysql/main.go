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
	host := viperConfigVariable("DB.HOST")
	port := viperConfigVariable("DB.PORT")
	user := viperConfigVariable("DB.USER")
	password := viperConfigVariable("DB.PASSWORD")
	name := viperConfigVariable("DB.NAME")
	fmt.Printf("host: " + host + "\n" +
		"port: " + port + "\n" +
		"user: " + user + "\n" +
		"password: " + password + "\n" +
		"name: " + name + "\n")
}
