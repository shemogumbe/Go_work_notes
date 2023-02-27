package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "RENNAISSANCE", Artist: "Beyonce", Price: 99.99},
	{ID: "2", Title: "MOTOMAMI", Artist: "Rosalia", Price: 17.99},
	{ID: "3", Title: "I Love You Jennifer B", Artist: "Jockstrap", Price: 39.99},
	{ID: "4", Title: "LP3", Artist: "Hippo Campus", Price: 9.99},
	{ID: "5", Title: "The Hardest Part", Artist: "Noah Cyrus", Price: 8.45},
	{ID: "6", Title: "If My Wife Knew I would be Dead", Artist: "CMAT", Price: 5.0},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")
	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}
func PostAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func main() {
	router := gin.Default()
	router.GET("albums", GetAlbums)
	router.GET("albums/:id", GetAlbumById)
	router.POST("albums", PostAlbums)

	router.Run("localhost:8080")
}
