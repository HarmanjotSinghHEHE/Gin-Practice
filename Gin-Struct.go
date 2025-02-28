package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Harman", Email: "harmans555@xmail.com"},
	{ID: 2, Name: "Shaina", Email: "shainas555@xmail.com"},
}

func main() {
	router := gin.Default()

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, users)
	})
	// to get specific users by ID
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, user := range users {
			if fmt.Sprintf("%d", user.ID) == id {
				c.JSON(http.StatusOK, user)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	})
	//to add a new user
	router.POST("/users", func(c *gin.Context) {
		var newUser User

		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		newUser.ID = len(users) + 1
		users = append(users, newUser)
		c.JSON(http.StatusCreated, newUser)
	})
	//To delete a user by ID
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, user := range users {
			if fmt.Sprintf("%d", user.ID) == id {
				users = append(users[:i], users[i+1])
				c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	})

	router.Run(":6969")

}
