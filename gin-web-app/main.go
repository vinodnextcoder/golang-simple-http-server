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

// / getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
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
	router.GET("/userData/:id", getUserByID)

	router.Run(":3000")
}
