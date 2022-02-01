package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthcheck", HealthCheck)
	router.Run("localhost:8080")

}

func HealthCheck(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"message": "health check is healthy"})
}

/*
Fix Git on local machine
mod and sum
API that is just health check
Test for functin
Logging
Documentation
Readme
Docker
GIthub
git actions
*/
