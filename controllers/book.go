package controllers

import (
	"book-management/database"
	"book-management/repository"
	"book-management/structs"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllBooks retrieves all books from the database
func GetAllBooks(c *gin.Context) {
	var result gin.H

	books, err := repository.GetAllBook(database.DbConnection)

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

// GetBook retrieves a single book from the database
func GetBook(c *gin.Context) {
	var result gin.H
	id, _ := strconv.Atoi(c.Param("id"))

	book, err := repository.GetBook(database.DbConnection, id)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": book,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertBook inserts a new book into the database
func InsertBook(c *gin.Context) {
	var book structs.Book

	// Bind JSON input ke struct Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Validasi tahun rilis
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Release year must be between 1980 and 2024"})
		return
	}

	// Konversi kolom thickness berdasarkan total_page
	if book.TotalPage > 100 {
		book.Thickness = "tebal" // Tebal jika halaman lebih dari 100
	} else {
		book.Thickness = "tipis" // Tipis jika halaman kurang dari 100
	}

	// Insert buku ke database
	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates an existing book in the database
func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	bookID, _ := strconv.Atoi(id)

	// Cek apakah buku dengan ID yang diberikan ada di database
	var existingBook structs.Book
	err := repository.FindBookByID(database.DbConnection, bookID, &existingBook)
	if err != nil {
		if err == sql.ErrNoRows {
			// Jika buku tidak ditemukan, kirimkan pesan error
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error_status":  "Data Not Found",
				"error_message": fmt.Sprintf("Book with id %v not found", id),
			})
			return
		}
		// Jika error lain terjadi, kirimkan pesan error umum
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}

	// Bind data JSON yang dikirimkan ke struct book
	var book structs.Book
	err = c.BindJSON(&book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Bad Request",
			"error_message": "Invalid JSON data",
		})
		return
	}

	// Validasi release_year (minimal 1980 dan maksimal 2024)
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2024 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{

			"error_message": "Release year must be between 1980 and 2024",
		})
		return
	}

	// Set ID buku yang akan diupdate
	book.ID = bookID

	// Lakukan update buku di database
	err = repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Internal Server Error",
			"error_message": err.Error(),
		})
		return
	}

	// Berikan respon sukses jika update berhasil
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfully updated", id),
		"book":    book,
	})
}

// DeleteBook deletes a book from the database
func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id
	err := repository.DeleteBook(database.DbConnection, book)

	// Cek apakah ada error saat menghapus buku
	if err != nil {
		// Jika buku tidak ditemukan, berikan pesan error
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			// Jika error lainnya, berikan pesan error
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		}
		return
	}

	// Berikan respon sukses jika penghapusan berhasil
	c.
		JSON(http.StatusOK, gin.H{"message": "Book deleted successfully", "book": book})
}
