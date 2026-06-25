package handlers

import (
	"blog-golang/models"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func render(w http.ResponseWriter, tmpl string, data any) {
	t, err := template.ParseFiles("templates/layout.html", tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.ExecuteTemplate(w, "layout", data)
	if err != nil {
		log.Printf("Erro ao executar o template %s: %v", tmpl, err)
	}
}

func (h *Handler) PostIndex(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetPublishedPosts(h.db)
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	render(w, "templates/posts/index.html", posts)
}

func (h *Handler) PostShow(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, err := models.GetPostBySlug(h.db, slug)
	if err == sql.ErrNoRows {
		http.Error(w, "Post não encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	render(w, "templates/posts/show.html", post)
}

func (h *Handler) PostsByCategory(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	posts, err := models.GetPostsByCategory(h.db, slug)
	if err == sql.ErrNoRows {
		http.Error(w, "Post não encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	render(w, "templates/categories/index.html", posts)
}

func (h *Handler) PostsByTag(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	posts, err := models.GetPostsByTag(h.db, slug)
	if err == sql.ErrNoRows {
		http.Error(w, "Post não encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	render(w, "templates/tags/index.html", posts)
}

func (h *Handler) AdminPostIndex(w http.ResponseWriter, r *http.Request) {
	posts, err := models.GetAllPostsAdmin(h.db)
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	render(w, "templates/posts/admin.html", posts)
}

// Get
type CreateTemplateData struct {
	Categories []models.Category
	Tags       []models.Tag
}

func (h *Handler) AdminPostNew(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetAllCategories(h.db)
	if err != nil {
		http.Error(w, "Erro ao carregar categorias", http.StatusInternalServerError)
		return
	}
	tags, err := models.GetAllTags(h.db)
	if err != nil {
		http.Error(w, "Erro ao carregar tags", http.StatusInternalServerError)
		return
	}
	viewData := CreateTemplateData{
		Categories: categories,
		Tags: tags,
	}
	render(w, "templates/posts/new.html", viewData)
}
