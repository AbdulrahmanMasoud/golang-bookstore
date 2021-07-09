package controllers

import (
	"net/http"

	"github.com/AbdulrahmanMasoud/Day1/models"
	"github.com/gin-gonic/gin"
)


func FindBooks(c *gin.Context){
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data":books})
}

func CreateBook(c *gin.Context){
	// Validation
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err !=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}

	//Create Book

	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)
	c.JSON(http.StatusOK,gin.H{"data":book})
}

func FindBook(c *gin.Context){
	var book models.Book

	if err := models.DB.Where("id=?",c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": book})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": book})
}


func UpdateBook(c *gin.Context){
	var book models.Book

	if err := models.DB.Where("id=?",c.Param("id")).First(&book).Error; err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found"})
		return
	}
	//Validation
	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	models.DB.Model(&book).Update(input)

	c.JSON(http.StatusOK,gin.H{"date":book})
}


func DeleteBook(c *gin.Context){
	var book models.Book
	if err := models.DB.Where("id=?",c.Param("id")).First(&book).Error; err !=nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error":"Book Not Found"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK,gin.H{"data":"Deleted Done"})
}


