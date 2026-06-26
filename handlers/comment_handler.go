package handlers

import (
	"blog-golang/models"
	"database/sql"
	"net/http"
	"strconv"
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

func (h *Handler) AdminCommentIndex(w http.ResponseWriter, r *http.Request) {
	comments, err := models.GetPendingComments(h.db)
	if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}
	render(w, "templates/comments/admin.html", comments)
}

func (h *Handler) AdminCommentApprove(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	err = models.ApproveComment(h.db, id)
	if err != nil {
		http.Error(w, "Erro ao aprovar comentário", http.StatusInternalServerError)
		return
	}


	http.Redirect(w, r, "/admin/comments", http.StatusSeeOther)
}

func(h *Handler) AdminCommentDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Erro ao delete comentário", http.StatusBadRequest)
		return
	}
	err = models.DeleteComment(h.db, id)
	if err != nil {
		http.Error(w, "Erro ao deletar comentário", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/admin/comments", http.StatusSeeOther)
}
