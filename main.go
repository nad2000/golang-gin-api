package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	log "github.com/sirupsen/logrus"
)

// Joke contains invormattion about a single Joke
type Joke struct {
	ID    int    `json:"id"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke"`
}

var db *gorm.DB

func main() {
	var err error
	db, err := gorm.Open("sqlite3", "jokes.db")

	if err != nil {
		log.WithError(err).Panicln("Failed to connect to the DB.")
	}
	defer db.Close()

	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		api.GET("/jokes", JokeHandler)
		api.POST("/jokes/like/:jokeID", LikeJoke)
	}

	// Start and run the server
	router.Run(":3000")
}

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
	c.Header("Context-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "JokeHandler not implemented yet",
	})
}

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
	c.Header("Context-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "LikeJoke not implemented yet",
	})
}
