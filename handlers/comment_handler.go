package handlers

import (
	"blog-golang/models"
	"net/http"
	"strconv"
)

func(h *Handler) CommentCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	slug := r.PathValue("slug")
	post, err := models.GetPostBySlug(h.db, slug)
	if err != nil {
		http.Error(w, "Id inválido", http.StatusBadRequest)
		return
	}
}
