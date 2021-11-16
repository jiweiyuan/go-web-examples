package main

import (
	"database/sql"
	"encoding/json"
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

	rows, err := db.Query("SELECT id, artist, title, price FROM album WHERE artist = ?", artist)
	if err != nil {
		return nil, fmt.Errorf("ListAlbumsByArtist %q: %v", artist, err)
	}
	defer rows.Close()

	for rows.Next() {
		var album Album
		if err := rows.Scan(&album.ID, &album.Artist, &album.Title, &album.Price); err != nil {
			return nil, fmt.Errorf("ListAlbumsByArtist %q: %v", artist, err)
		}
		albums = append(albums, album)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("ListAlbumsByArtist %q: %v", artist, err)
	}
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

// CreateAlbum add the specified album to the database,
// return the album ID of the new album
func CreateAlbum(album *Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (artist, title, price) VALUES (?, ?, ?)", album.Artist, album.Price, album.Price)
	if err != nil {
		return 0, fmt.Errorf("CreateAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateAlbum: %v", err)
	}
	return id, nil

}

func UpdateAlbum(album *Album) {

}

func DeleteAlbum(id int64)(bool, error) {
	result, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return false, fmt.Errorf("DeleteAlbum: %v", err)
	}
	check, err := result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("DeleteAlbum: %v", err)
	}

	if check == 1 {
		return true, nil
	} else {
		return false, fmt.Errorf("no item to delete")
	}

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
	s, _ := json.Marshal(albumOne)
	fmt.Printf("Get album of 2: %v\n", string(s))

	albumsX, _ := ListAlbumsByArtist("徐悲鸿")
	for i, album := range albumsX {
		s, _ := json.Marshal(album)
		fmt.Printf("ListAlbumsByArtist = 徐悲鸿 album %v: %v\n", i, string(s))
	}

	albId, err := CreateAlbum(&Album{
		Title: "The Modern Album",
		Artist: "berries",
		Price: 23.79,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album:%v\n", albId)

	resultNew, _ := DeleteAlbum(albId)
	fmt.Printf("DeleteAlbum: %v\n", resultNew)
	result100, _ := DeleteAlbum(100)
	fmt.Printf("DeleteAlbum: %v\n", result100)

}
