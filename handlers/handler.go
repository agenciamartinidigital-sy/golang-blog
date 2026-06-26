package handlers

import (
	"blog-golang/models"
	"database/sql"
	"time"
)

type Handler struct {
	db *sql.DB
}

type CreateTemplateData struct {
	Categories []models.Category
	Tags       []models.Tag
}

type AdminPostEditData struct {
	Post       models.Post
	Categories []models.Category
	Tags       []models.Tag
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
