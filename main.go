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

func setUpDB() {
	te := db.HasTable(&Joke{})
	db.AutoMigrate(&Joke{})
	if !te {
		for _, j := range []Joke{
			{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
			{2, 0, "What do you call a fake noodle? An Impasta."},
			{3, 0, "How many apples grow on a tree? All of them."},
			{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
			{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
			{6, 0, "Why did the coffee file a police report? It got mugged."},
			{7, 0, "How does a penguin build it's house? Igloos it together."},
		} {
			if err := db.Create(&j).Error; err != nil {
				log.WithError(err).Errorln("Failed to insert a joke", j)
			}
		}

	}
}

func main() {
	var err error
	db, err = gorm.Open("sqlite3", "jokes.db")

	if err != nil {
		log.WithError(err).Panicln("Failed to connect to the DB.")
	}
	defer db.Close()
	setUpDB()

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
