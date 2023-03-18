package main

import (
	"github.com/gin-gonic/gin"
	"hilarick.com/goinapi/controllers"
	"hilarick.com/goinapi/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run(":3000")
}
