package controllers

import (
	"book-management/database"
	"book-management/repository"
	"book-management/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllCategories retrieves all categories from the database
func GetAllCategories(c *gin.Context) {
	var result gin.H

	categories, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

// GetBooksByCategory retrieves all books by category from the database
func GetBooksByCategory(c *gin.Context) {
	var result gin.H
	id, _ := strconv.Atoi(c.Param("id"))

	books, err := repository.GetBooksByCategory(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

// GetCategory retrieves a single category from the database
func GetCategory(c *gin.Context) {
	var result gin.H
	id, _ := strconv.Atoi(c.Param("id"))

	category, err := repository.GetCategory(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": category,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertCategory inserts a new category into the database
func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates an existing category in the database
func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	// Bind JSON input to category struct
	err := c.BindJSON(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
		})
		return
	}

	category.ID = id

	// Check if the category exists before attempting to update
	err = repository.GetCategoryByID(database.DbConnection, &category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
		return
	}

	// Proceed to update the category if it exists
	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Category updated successfully",
		"category": category,
	})
}

// DeleteCategory deletes a category from the database
func DeleteCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id

	// Check if the category exists before attempting to delete
	err := repository.GetCategoryByID(database.DbConnection, &category)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Category not found",
		})
		return
	}

	// Proceed to delete the category if it exists
	err = repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete category",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category deleted successfully",
	})
}
