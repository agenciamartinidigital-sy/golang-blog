package handlers

import (
	"blog-golang/models"
	"database/sql"
	"net/http"
)

func (h *Handler) CommentCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	authorName := r.FormValue("author_name")
	authorEmail := r.FormValue("author_email")
	content := r.FormValue("content")

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
	err = models.CreateComment(h.db, post.ID, authorName, authorEmail, content)
	if err != nil {
		http.Error(w, "Erro ao criar comentário", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/posts/"+slug, http.StatusSeeOther)
}

func(h *Handler) AdminCommentIndex(w http.ResponseWriter, r *http.Request) {
	post, err := models.GetPendingComments(h.db)
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	render(w, "templates/comments/admin.html", post)
}

func(h *Handler) AdminCommentAprove(w http.ResponseWriter, r http.Request) {

	

	http.Redirect(w, r, "/admin/comments", http.StatusSeeOther)
}
