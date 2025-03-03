package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// now i will define the struct along with the json names

type ToDo struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Completetion bool   `json:"completetion"`
}

var todos []ToDo

func getTodos(c *gin.Context) {
	if len(todos) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Task DIkhane ko Nhi hai",
			"todos":   []ToDo{},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}

func getTodobyID(c *gin.Context) {
	id := c.Param("id")
	for _, task := range todos {
		if fmt.Sprintf("%d", task.ID) == id {
			c.JSON(http.StatusOK, gin.H{
				"todo": task,
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"Error": "Task nahi Hai!!!!1",
	})
}

func createToDo(c *gin.Context) {
	var newtodo ToDo
	//code if the JSON data is in invaid format or is invalid
	if err := c.ShouldBindJSON(&newtodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Input",
		})
		return
	}

	if newtodo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Title can not be empty",
		})
	}

	newtodo.ID = len(todos) + 1
	newtodo.Completetion = false
	todos = append(todos, newtodo)
	c.JSON(http.StatusOK, gin.H{
		"message": "Task added Successfully",
		"todo":    newtodo,
	})
}

func updateTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range todos {
		if fmt.Sprintf("%d", task.ID) == id {
			todos[i].Completetion = true
			c.JSON(http.StatusOK, gin.H{
				"message": "Task marked as Completed",
				"todo":    todos[i],
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"error": "Task Not Found",
	})
}

func deleteTask(c *gin.Context) {
	id := c.Param("id")
	for i, task := range todos {
		if fmt.Sprintf("%d", task.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Task deleted Successfully",
				//"todo":    todos[i],
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Task not Found",
	})
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodobyID)
	router.POST("/todos", createToDo)
	router.PUT("/todos/:id", updateTask)
	router.DELETE("todos/:id", deleteTask)
	router.Run(":6969")
}
