package models

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID          int       `json:"id"`
	PostID      int       `json:"post_id"`
	PostTitle   string    `json:"post_title"`
	PostSlug    string    `json:"post_slug"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	Content     string    `json:"content"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type Post struct {
	ID          int          `json:"id"`
	CategoryID  int          `json:"category_id"`
	Title       string       `json:"title"`
	Slug        string       `json:"slug"`
	Content     string       `json:"content"`
	Status      string       `json:"status"`
	PublishedAt sql.NullTime `json:"published_at"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Tags        []Tag        `json:"tags"`
	Comments    []Comment    `json:"comments"`
	Category    Category     `json:"category"`
}

type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
