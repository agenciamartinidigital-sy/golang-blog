package handlers

import (
	"blog-golang/models"
	"blog-golang/utils"
	"database/sql"
	"net/http"
)

// POST
func (h *Handler) CategoryCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	slug := utils.Slugify(name)
	err := models.CreateCategory(h.db, name, slug)
	if err != nil {
		http.Error(w, "Categoria não criada", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}

// GET
func (h *Handler) CategoryIndex(w http.ResponseWriter, r *http.Request) {
	category, err := models.GetAllCategories(h.db)
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	render(w, "templates/categories/index.html", category)
}

// POST
func (h *Handler) CategoryUpdate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	slug := r.PathValue("slug")
	if slug == "" {
		http.Error(w, "Slug não fornecido", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name não fornecido", http.StatusBadRequest)
		return
	}

	newSlug := utils.Slugify(name)

	category, err := models.GetCategoryBySlug(h.db, slug)
	if err == sql.ErrNoRows{
		http.Error(w, "Categoria não encontrada", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	err = models.UpdateCategory(h.db, category.ID, name, newSlug)
	if err != nil {
		http.Error(w, "Erro ao atualizar categoria", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/categories", http.StatusSeeOther)
}


