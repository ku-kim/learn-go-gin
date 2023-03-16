package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// ex : https://gin-gonic.com/zh-tw/docs/examples/using-middleware/

func main() {
	router := gin.New()
	router.Use(gin.Logger()) // 전역 미들웨어 추가
	router.Use(gin.Recovery())
	//router := gin.Default() // Default()는 사실 gin.New()에 기본이 되는 미들웨어 추가함

	v1API := router.Group("/v1")
	v1API.GET("/albums", CustomLoggerMiddleware(), getAlbums) // 특정 메서드 미들웨어 추가 가능
	v1API.GET("/albums/:id", getAlbumByID)
	v1API.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
