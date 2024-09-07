package repository

import (
	"book-management/structs"
	"database/sql"
)

// GetAllUser retrieves all users from the database.
func GetAllUser(db *sql.DB) (result []structs.User, err error) {
	query := "SELECT * FROM users"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user structs.User

		err = rows.Scan(
			&user.ID,         // BIGINT -> int
			&user.Username,   // varchar(256) -> string
			&user.Password,   // varchar(256) -> string
			&user.CreatedAt,  // TIMESTAMP -> time.Time
			&user.CreatedBy,  // varchar(256) -> string
			&user.ModifiedAt, // TIMESTAMP -> time.Time
			&user.ModifiedBy, // varchar(256) -> string
		)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}

	return result, nil
}

// InsertUser inserts a new user into the database.
func InsertUser(db *sql.DB, user structs.User) (err error) {
	query := `
		INSERT INTO users (
			id, username, password, created_at, created_by, modified_at, modified_by
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err = db.Exec(query, user.ID, user.Username, user.Password, user.CreatedAt, user.CreatedBy, user.ModifiedAt, user.ModifiedBy)
	return err
}

// UpdateUser updates an existing user in the database.
func UpdateUser(db *sql.DB, user structs.User) (err error) {
	query := `
		UPDATE users
		SET 
			username = $1, 
			password = $2, 
			created_at = $3, 
			created_by = $4, 
			modified_at = $5, 
			modified_by = $6
		WHERE id = $7
	`

	_, err = db.Exec(query, user.Username, user.Password, user.CreatedAt, user.CreatedBy, user.ModifiedAt, user.ModifiedBy, user.ID)
	return err
}

// DeleteUser deletes a user from the database.
func DeleteUser(db *sql.DB, user structs.User) (err error) {
	query := "DELETE FROM users WHERE id = $1"

	_, err = db.Exec(query, user.ID)
	return err
}

// GetUserByUsername retrieves a user by username from the database.
func GetUserByUsername(db *sql.DB, username string) (user structs.User, err error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err = db.QueryRow(query, username).Scan(
		&user.ID,       // BIGINT -> int
		&user.Username, // varchar(256) -> string
		&user.Password, // varchar(256) -> string
	)
	return user, err
}
