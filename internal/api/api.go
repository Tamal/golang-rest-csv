package api

import (
	"fmt"
	"net/http"

	"github.com/emp/internal/db"
	"github.com/emp/internal/db/model"
	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine {
	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/users/:id", getUser)
	router.POST("/users", addUser)
	router.DELETE("/users/:id", deleteUser)

	return router
}

func getUsers(c *gin.Context) {

	svc := db.UserDataService{FileName: "emp.csv"}
	users, err := svc.GetUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")

	svc := db.UserDataService{FileName: "emp.csv"}
	user, err := svc.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func addUser(c *gin.Context) {
	svc := db.UserDataService{FileName: "emp.csv"}

	var user model.User
	c.BindJSON(&user)
	fmt.Println("User", user)
	users, err := svc.AddUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	svc := db.UserDataService{FileName: "emp.csv"}
	user, err := svc.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
