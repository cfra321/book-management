package controllers

import (
	"book-management/database"
	"book-management/repository"
	"book-management/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers(c *gin.Context) {
	var result gin.H

	users, err := repository.GetAllUser(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": users,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertUser inserts a new user into the database
func InsertUser(c *gin.Context) {
	var user structs.User

	// Bind JSON input to User struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Insert the user into the database
	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not insert user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

// UpdateUser updates an existing user in the database
func UpdateUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	// Bind JSON input to User struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Set the ID from the URL parameter
	user.ID = id

	// Check if password is provided for update and hash it
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
			return
		}
		user.Password = string(hashedPassword)
	}

	// Update the user in the database
	err := repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// DeleteUser deletes a user from the database
func DeleteUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("id"))

	user.ID = id
	err := repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, user)
}
