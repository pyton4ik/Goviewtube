package app

import (
	"github.com/pyton4ik/Goviewtube/goserver/internal/transport/routes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {
		// Get the client IP address
		clientIP := c.ClientIP()

		// Get the current time
		now := time.Now()
		// Log the request
		log.Printf("[%s] %s %s %s", now.Format(time.RFC3339), c.Request.Method, c.Request.URL.Path, clientIP)

		// Proceed to the next handler
		c.Next()
	}
}

func Run() {
	router := gin.Default()
	// Simple group: v1
	api := router.Group("/api")
	routes.SetRoutes(api)
	router.Run("0.0.0.0:8067") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
