package repository

import (
	"book-management/structs"
	"database/sql"
	"errors"
)

func GetAllCategory(db *sql.DB) (result []structs.Category, err error) {
	sql := "SELECT * FROM category"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var category structs.Category

		err = rows.Scan(
			&category.ID,
			&category.Name,
			&category.CreatedAt,
			&category.CreatedBy,
			&category.ModifiedAt,
			&category.ModifiedBy)
		if err != nil {
			return
		}

		result = append(result, category)
	}

	return
}

func GetBooksByCategory(db *sql.DB, id int) (result []structs.Book, err error) {
	// Specify the columns to avoid issues with column counts
	sql := "SELECT id, title, category_id, created_at, created_by, modified_at, modified_by FROM book WHERE category_id = $1"

	rows, err := db.Query(sql, id)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book structs.Book

		// Map the selected columns to the struct fields
		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.CategoryID,
			&book.CreatedAt,
			&book.CreatedBy,
			&book.ModifiedAt,
			&book.ModifiedBy)
		if err != nil {
			return
		}

		result = append(result, book)
	}

	return
}

func GetCategory(db *sql.DB, id int) (category structs.Category, err error) {
	sql := "SELECT * FROM category WHERE id = $1"

	err = db.QueryRow(sql, id).Scan(
		&category.ID,
		&category.Name,
		&category.CreatedAt,
		&category.CreatedBy,
		&category.ModifiedAt,
		&category.ModifiedBy)

	return
}

func InsertCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "INSERT INTO category(id, name, created_at, created_by, modified_at, modified_by) VALUES ($1, $2, $3, $4, $5, $6)"

	errs := db.QueryRow(sql, category.ID, category.Name, category.CreatedAt, category.CreatedBy, category.ModifiedAt, category.ModifiedBy)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "UPDATE category SET name = $1, created_at = $2, created_by = $3, modified_at = $4, modified_by = $5 WHERE id = $6"

	errs := db.QueryRow(sql, category.Name, category.CreatedAt, category.CreatedBy, category.ModifiedAt, category.ModifiedBy, category.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, category structs.Category) (err error) {
	sql := "DELETE FROM category WHERE id = $1"

	errs := db.QueryRow(sql, category.ID)
	return errs.Err()
}

// GetCategoryByID retrieves a category from the database by its ID
func GetCategoryByID(db *sql.DB, category *structs.Category) error {
	query := "SELECT id, name FROM categories WHERE id = ?"
	err := db.QueryRow(query, category.ID).Scan(&category.ID, &category.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("category not found")
		}
		return err
	}

	return nil
}
