package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type userDetail struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

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

var users = []userDetail{
	{ID: "1", Name: "Blue Train", Age: 56},
	{ID: "2", Name: "Akash", Age: 11},
	{ID: "3", Name: "pradip", Age: 56},
}

func getUsers(c *gin.Context) {

	c.JSON(http.StatusOK, users)
}

// postAlbums adds an album from JSON received in the request body.
func addUser(c *gin.Context) {

	var newUser userDetail

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new album to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

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
	router.GET("/user", getUsers)
	router.POST("/user", addUser)

	router.Run(":3000")
}
