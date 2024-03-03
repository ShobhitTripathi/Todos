// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var todos []Todo

// Todo structure
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// GetTodos returns the list of todos
func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

// GetTodoByID returns a specific todo by ID
func GetTodoByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// Find the todo with the specified ID
	for _, todo := range todos {
		if id == todo.ID {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// CreateTodo creates a new todo
func CreateTodo(c *gin.Context) {
	var newTodo Todo
	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a unique ID
	newTodo.ID = len(todos) + 1

	// Add the new todo to the list
	todos = append(todos, newTodo)

	c.JSON(http.StatusCreated, newTodo)
}

// UpdateTodo updates an existing todo
func UpdateTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updatedTodo Todo
	if err := c.ShouldBindJSON(&updatedTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the todo with the specified ID
	for i, todo := range todos {
		if id == todo.ID {
			// Update the todo
			todos[i] = updatedTodo
			c.JSON(http.StatusOK, updatedTodo)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

// DeleteTodo deletes a todo by ID
func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// Find the todo with the specified ID
	for i, todo := range todos {
		if id == todo.ID {
			// Remove the todo from the list
			todos = append(todos[:i], todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func main() {
	r := gin.Default()

	// Define API routes
	r.GET("/todos", GetTodos)
	r.GET("/todos/:id", GetTodoByID)
	r.POST("/todos", CreateTodo)
	r.PUT("/todos/:id", UpdateTodo)
	r.DELETE("/todos/:id", DeleteTodo)

	// Run the server on port 8080
	r.Run(":8080")
}
