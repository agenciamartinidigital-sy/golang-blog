package handlers

import (
	"blog-golang/models"
	"database/sql"
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
