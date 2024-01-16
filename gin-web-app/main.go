package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vinodnextcoder/go-web-app/docs"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample gin web server

// @contact.name   vinod
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:3001
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	router := gin.Default()

	router.GET("/", helloCall)

	router.GET("/about", aboutCall)
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

// helloCall godoc
// @Summary hellow example
// @Schemes
// @Description Hello
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Hello, You created a Web App!
// @Router / [get]
func helloCall(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello, You created a Web App!"})
}

// aboutCall godoc
// @Summary about routes
// @Schemes
// @Description about
// @Tags about
// @Accept json
// @Produce json
// @Success 200 {string} About Us
// @Router /about [get]
func aboutCall(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "About Us"})
}

// @Summary get user name
// @Accept  json
// @Produce  json
// @Param        name   path      string  true  "usernmae"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user/{name} [get]
func getUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
}

var users = []userDetail{
	{ID: "1", Name: "Blue Train", Age: 56},
	{ID: "2", Name: "Akash", Age: 11},
	{ID: "3", Name: "pradip", Age: 56},
}

// @Summary users list
// @Description users list
// @Accept  json
// @Produce  json
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user [get]
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

// @Summary get user id
// @Description user get  by id
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Param        id   path      int  true  "userid ID"
// @Success      200
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /user/{id} [get]
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
