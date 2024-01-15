package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	router := gin.Default()

	router.GET("/", helloCall)

	router.GET("/about", aboutCall)
	router.GET("/contact", contactCall)
	router.GET("/user/:name", getUser)
	router.GET("/user", getUsers)
	router.POST("/user", addUser)
	router.GET("/userData/:id", getUserByID)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":3001")
}

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

func addUser(c *gin.Context) {

	var newUser userDetail

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func getUserByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range users {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
