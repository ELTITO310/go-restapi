package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sonyarouje/simdb"
)

type Album struct {
	AlbumID string  `json:"albumid"`
	Title   string  `json:"title"`
	Artist  string  `json:"artist"`
	Price   float32 `json:"price"`
}

func (c Album) ID() (jsonField string, value interface{}) {
	value = c.ID
	jsonField = "albumid"
	return
}

func main() {

	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:3000")

	fmt.Println("SERVER RUNNING")

}

func getAlbums(c *gin.Context) {

	driver, err := db.New("data")

	if err != nil {
		log.Fatal(err)
	}

	dataAll := driver.Open(Album{})
	datas := dataAll.Get().RawArray()

	fmt.Println(datas)

	c.IndentedJSON(http.StatusOK, datas)

}

func postAlbums(c *gin.Context) {
	driver, err := db.New("data")
	if err != nil {
		log.Fatal(err)
	}
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	driver.Insert(newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	driver, err := db.New("data")

	if err != nil {
		log.Fatal(err)
	}

	dataAll := driver.Open(Album{}).Where("albumid", "=", id).Get().Raw()

	c.IndentedJSON(http.StatusOK, dataAll)

}
