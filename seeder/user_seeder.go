package seeder

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// SeedUsers seeds the database with initial user data
func SeedUsers(db *sql.DB) {
	// Data pengguna yang akan di-seed
	users := []struct {
		ID         int
		Username   string
		Password   string
		CreatedAt  string
		CreatedBy  string
		ModifiedAt string
		ModifiedBy string
	}{
		{
			ID:         1,
			Username:   "admin",
			Password:   "password", // ini akan di-hash
			CreatedAt:  "2024-08-03T14:55:00Z",
			CreatedBy:  "admin",
			ModifiedAt: "2024-08-03T15:00:00Z",
			ModifiedBy: "admin",
		},
	}

	for _, user := range users {
		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Tidak dapat meng-hash password: %v", err)
		}

		// Periksa apakah pengguna sudah ada
		var existingID int
		err = db.QueryRow(`SELECT id FROM users WHERE username = $1`, user.Username).Scan(&existingID)
		if err != nil && err != sql.ErrNoRows {
			log.Fatalf("Error memeriksa apakah pengguna ada: %v", err)
		}

		// Jika pengguna tidak ada, lakukan insert
		if err == sql.ErrNoRows {
			_, err = db.Exec(`
				INSERT INTO users (id, username, password, created_at, created_by, modified_at, modified_by)
				VALUES ($1, $2, $3, $4, $5, $6, $7)
			`, user.ID, user.Username, hashedPassword, user.CreatedAt, user.CreatedBy, user.ModifiedAt, user.ModifiedBy)

			if err != nil {
				log.Fatalf("Tidak dapat memasukkan pengguna: %v", err)
			}
		}
	}

	log.Println("Pengguna berhasil di-seed.")
}
