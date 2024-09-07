package repository

import (
	"book-management/structs"
	"database/sql"
	"errors"
)

func GetAllBook(db *sql.DB) (result []structs.Book, err error) {
	sql := "SELECT * FROM book"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book structs.Book

		err = rows.Scan(
			&book.ID,          // id
			&book.Title,       // title
			&book.Description, // description
			&book.ImageURL,    // image_url
			&book.ReleaseYear, // release_year
			&book.Price,       // price
			&book.TotalPage,   // total_page
			&book.Thickness,   // thickness
			&book.CategoryID,  // category_id
			&book.CreatedAt,   // created_at
			&book.CreatedBy,   // created_by
			&book.ModifiedAt,  // modified_at
			&book.ModifiedBy,  // modified_by
		)
		if err != nil {
			return
		}

		result = append(result, book)
	}

	return
}

func GetBook(db *sql.DB, id int) (book structs.Book, err error) {
	sql := "SELECT * FROM book WHERE id = $1"

	err = db.QueryRow(sql, id).Scan(
		&book.ID,          // id
		&book.Title,       // title
		&book.Description, // description
		&book.ImageURL,    // image_url
		&book.ReleaseYear, // release_year
		&book.Price,       // price
		&book.TotalPage,   // total_page
		&book.Thickness,   // thickness
		&book.CategoryID,  // category_id
		&book.CreatedAt,   // created_at
		&book.CreatedBy,   // created_by
		&book.ModifiedAt,  // modified_at
		&book.ModifiedBy,  // modified_by
	)

	return
}

func InsertBook(db *sql.DB, book structs.Book) error {
	sql := `
        INSERT INTO book (id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by) 
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	_, err := db.Exec(sql,
		book.ID,
		book.Title,
		book.Description,
		book.ImageURL,
		book.ReleaseYear,
		book.Price,
		book.TotalPage,
		book.Thickness,
		book.CategoryID,
		book.CreatedAt,
		book.CreatedBy,
		book.ModifiedAt,
		book.ModifiedBy,
	)

	if err != nil {
		return err
	}

	return nil
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	sql := "UPDATE book SET title = $1, description = $2, image_url = $3, release_year = $4, price = $5, total_page = $6, thickness = $7, category_id = $8, created_at = $9, created_by = $10, modified_at = $11, modified_by = $12 WHERE id = $13"

	errs := db.QueryRow(sql, book.Title, book.Description, book.ImageURL, book.ReleaseYear, book.Price, book.TotalPage, book.Thickness, book.CategoryID, book.CreatedAt, book.CreatedBy, book.ModifiedAt, book.ModifiedBy, book.ID)

	return errs.Err()
}

func FindBookByID(db *sql.DB, id int, book *structs.Book) error {
	sql := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM book WHERE id = $1"
	return db.QueryRow(sql, id).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	sql := "DELETE FROM book WHERE id = $1"

	// Menggunakan Exec karena QueryRow tidak cocok untuk operasi DELETE
	result, err := db.Exec(sql, book.ID)
	if err != nil {
		return err
	}

	// Cek jumlah baris yang terpengaruh (harus 1 jika berhasil menghapus)
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Jika tidak ada baris yang dihapus, return error untuk buku tidak ditemukan
	if rowsAffected == 0 {
		return errors.New("book not found")
	}

	return nil
}
