package handlers

import (
	"blog-golang/models"
	"blog-golang/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"time"
)

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

func (h *Handler) AdminPostNew(w http.ResponseWriter, r *http.Request) {
	categories, err := models.GetAllCategories(h.db)
	if err != nil {
		http.Error(w, "Erro ao carregar categories", http.StatusInternalServerError)
		return
	}
	tags, err := models.GetAllTags(h.db)
	if err != nil {
		http.Error(w, "Erro ao carregar tags", http.StatusInternalServerError)
		return
	}
	viewData := CreateTemplateData{
		Categories: categories,
		Tags:       tags,
	}
	render(w, "templates/posts/new.html", viewData)
}

func (h *Handler) AdminPostCreate(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	title := r.FormValue("title")
	content := r.FormValue("content")
	status := r.FormValue("status")

	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		http.Error(w, "Categoria inválida", http.StatusBadRequest)
		return
	}
	tagStrings := r.Form["tags"] // []string

	// cria o slice de inteiros
	tagIDs := make([]int, len(tagStrings))

	for i, tagStr := range tagStrings {
		id, err := strconv.Atoi(tagStr)
		if err != nil {
			http.Error(w, "Uma ou mais tags fornecidads são inválidas", http.StatusBadRequest)
			return
		}
		tagIDs[i] = id
	}
	slug := utils.Slugify(title)

	var publishedAt sql.NullTime
	if status == "published" {
		publishedAt = sql.NullTime{Time: time.Now(), Valid: true}
	}

	err = models.CreatePost(h.db, title, content, slug, status, categoryID, publishedAt, tagIDs)
	if err != nil {
		http.Error(w, "Erro ao salvar o post no banco de dados", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
}

func (h *Handler) AdminPostEdit(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "Slug não fornecido", http.StatusBadRequest)
		return
	}
	post, err := models.GetPostBySlug(h.db, slug)
	if err == sql.ErrNoRows {
		http.Error(w, "Post não encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	categories, err := models.GetAllCategories(h.db)
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	tags, err := models.GetAllTags(h.db)
	if err != nil {
		http.Error(w, "Erro ao carregar tags", http.StatusInternalServerError)
		return
	}
	data := AdminPostEditData{
		Post:       post,
		Categories: categories,
		Tags:       tags,
	}

	render(w, "templates/posts/edit.html", data)
}

func (h *Handler) AdminPostUpdate(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "Slug não fornecido", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Erro ao ler o formulário", http.StatusBadRequest)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	statusStr := r.FormValue("status")

	newSlug := utils.Slugify(title)

	categoryID, err := strconv.Atoi(r.FormValue("category_id"))
	if err != nil {
		http.Error(w, "Categoria inválida", http.StatusBadRequest)
		return
	}

	tagString := r.Form["tags"]
	tagIDs := make([]int, len(tagString))

	for i, tagStr := range tagString {
		id, err := strconv.Atoi(tagStr)
		if err != nil {
			http.Error(w, "Uma ou mais tags fornecidas são inválidas", http.StatusBadRequest)
			return
		}
		tagIDs[i] = id
	}

	post, err := models.GetPostBySlug(h.db, slug)
	if err == sql.ErrNoRows {
		http.Error(w, "Post não encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	var publishedAt sql.NullTime
	if statusStr == "published" {
		if post.PublishedAt.Valid {
			publishedAt = post.PublishedAt
		} else {
			publishedAt = sql.NullTime{Time: time.Now(), Valid: true}
		}
	}

	err = models.UpdatePost(h.db, title, content, newSlug, statusStr, post.ID, categoryID, publishedAt, tagIDs)
	if err != nil {
		http.Error(w, "Erro ao atualizar o post no banco de dados", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)
}

func (h *Handler) AdminPostDelete(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	post, err := models.GetPostBySlug(h.db, slug)
	if err == sql.ErrNoRows {
		http.Error(w, "Não encotrado", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Printf("AdminPostDelete: %v", err)
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}
	err = models.DeletePost(h.db, post.ID)
	if err != nil {
		http.Error(w, "Erro ao deletar", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/posts", http.StatusSeeOther)

}
