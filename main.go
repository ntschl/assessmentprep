package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()
	r.GET("/albums", getRandomAlbum)
	r.GET("/albums/:id", getAlbumByID)
	r.POST("/albums", postAlbum)
	r.Run("0.0.0.0:80")
}

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title,omitempty"`
	Artist string `json:"artist,omitempty"`
}

var albums = map[string]album{
	"c5bbc031-b0af-4709-9f40-9a050949c555": {"c5bbc031-b0af-4709-9f40-9a050949c555", "Lonerism", "Tame Impala"},
	"2d2dd8de-466f-4110-ade9-6fad94f60520": {"2d2dd8de-466f-4110-ade9-6fad94f60520", "Bloom", "Beach House"},
	"765212ed-65c3-4b34-876c-0b504fec271f": {"765212ed-65c3-4b34-876c-0b504fec271f", "Dedicated", "Carly Rae Jepsen"},
	"41467023-50a2-4de4-99f0-56224e92207f": {"41467023-50a2-4de4-99f0-56224e92207f", "Shrines", "Purity Ring"},
	"eb808439-b8c4-4a8c-9306-36a8ccb49727": {"eb808439-b8c4-4a8c-9306-36a8ccb49727", "Perfume", "Wand"},
}

func randomizer() string {
	var keys []string
	for key := range albums {
		keys = append(keys, key)
	}
	rando := rand.Intn(len(albums))
	return keys[rando]
}

func getRandomAlbum(c *gin.Context) {
	key := randomizer()
	album := albums[key]
	c.JSON(http.StatusOK, album)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album := albums[id]
	c.JSON(http.StatusOK, album)
}

func postAlbum(c *gin.Context) {
	id := uuid.New().String()
	var newAlbum album
	var returnAlbum album
	returnAlbum.ID = id
	newAlbum.ID = id
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums[id] = newAlbum
	c.JSON(http.StatusCreated, returnAlbum)
}
