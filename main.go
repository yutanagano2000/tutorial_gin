package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var albums = []album{
	{ID: "1", Title: "Specialz"},
	{ID: "2", Title: "IRIS OUT"},
	{ID: "3", Title: "Lemon"},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.Run("localhost:8080")
}
