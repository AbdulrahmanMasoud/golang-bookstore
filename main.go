package main

import (
// "net/http"
"github.com/gin-gonic/gin"

"github.com/AbdulrahmanMasoud/Day1/models"
"github.com/AbdulrahmanMasoud/Day1/controllers"
)
func main() {
	route := gin.Default()

	models.ConnectDataBase()
	route.GET("/books", controllers.FindBooks)
	route.POST("/books", controllers.CreateBook)
	route.GET("/book/:id", controllers.FindBook)
	route.PUT("/book/:id", controllers.UpdateBook)
	route.DELETE("/book/:id", controllers.DeleteBook)
	route.Run()
}
