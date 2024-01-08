package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloCall(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, You created a Web App!"})
}

func aboutCall(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "About Us"})
}
func contactCall(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Contact Us"})
}

func getUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
}

func main() {
	router := gin.Default()

	// Define a route handler for the root path
	router.GET("/", helloCall)

	// Define a route handler for "/about"
	router.GET("/about", aboutCall)

	// Define a route handler for "/contact"
	router.GET("/contact", contactCall)

	// Add this route handler
	// with param
	router.GET("/user/:name", getUser)

	router.Run(":3000")
}
