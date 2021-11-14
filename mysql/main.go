package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
)

var db *sql.DB

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
		log.Fatalf("Invalid type assertion")
	}

	return value
}

type Album struct {
	ID     int
	Artist string
	Title  string
	Price  float32
}

func ListAlbumsByArtist(artist string) ([]Album, error) {
	var albums []Album
	return albums, nil
}

func GetAlbumById(id int) (Album, error) {
	var a Album
	row := db.QueryRow("SELECT id, artist, title, price FROM album WHERE id = ?", id)
	if err := row.Scan(&a.ID, &a.Artist, &a.Title, &a.Price); err != nil {
		if err == sql.ErrNoRows {
			return a, fmt.Errorf("can find ablum of %v: no such ablum", id)
		}
		return a, fmt.Errorf("can't find ablum of %v: %s", id, err)
	}

	return a, nil
}

func CreateAlbum(album *Album) {

}

func UpdateAlbum(album *Album) {

}

func deleteAlbum(id int) {

}

func main() {
	config := mysql.Config{
		Passwd: viperConfigVariable("DB.PASSWORD"),
		User:   viperConfigVariable("DB.USER"),
		DBName: viperConfigVariable("DB.NAME"),
		Addr:   viperConfigVariable("DB.HOST") + ":" + viperConfigVariable("DB.PORT"),
		Net:    "tcp",
	}

	var err error
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config.FormatDSN())

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connect")
	}

	albumOne, _ := GetAlbumById(2)
	fmt.Printf("Get album of 2: %v", albumOne)
}
