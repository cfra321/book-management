package structs

import "time"

type Book struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
	ReleaseYear int       `json:"release_year,omitempty"`
	Price       int       `json:"price,omitempty"`
	TotalPage   int       `json:"total_page,omitempty"`
	Thickness   string    `json:"thickness,omitempty"`
	CategoryID  int       `json:"category_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	CreatedBy   string    `json:"created_by,omitempty"`
	ModifiedAt  time.Time `json:"modified_at,omitempty"`
	ModifiedBy  string    `json:"modified_by,omitempty"`
}

type Category struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`
}
