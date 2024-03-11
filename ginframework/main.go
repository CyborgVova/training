package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id        int64  `json:"id"`
	Author    string `json:"author"`
	AlbumName string `json:"albumname"`
	Year      int32  `json:"year"`
}

var (
	albums = []album{
		{Id: 1, Author: "Viktor Tsoy", AlbumName: "45", Year: 1989},
		{Id: 2, Author: "Abba", AlbumName: "Abba", Year: 1987},
		{Id: 3, Author: "Elvis Praisly", AlbumName: "Liriks", Year: 1982},
	}
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")
	num, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "argument is not a number"})
		return
	}

	for _, v := range albums {
		if v.Id == int64(num) {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var album album

	if err := c.BindJSON(&album); err != nil {
		return
	}

	albums = append(albums, album)

	c.IndentedJSON(http.StatusCreated, albums)
}

func main() {
	router := gin.Default()

	router.StaticFile("./favicon.ico", "./static/favicon.ico")
	router.LoadHTMLFiles("./index.html")
	router.SetTrustedProxies(nil)

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
