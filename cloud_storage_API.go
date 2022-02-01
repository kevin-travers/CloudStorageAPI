package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//main sets up gin router and create endpoints that can be called
func main() {
	router := gin.Default()
	router.GET("/healthcheck", HealthCheck)
	router.Run("localhost:8080")

}

//Health check is an api that returns a status ok and message to let call know that the api endpoint is up and running
func HealthCheck(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "health check is healthy"})
}
