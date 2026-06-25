package models

import (
	"database/sql"
	"fmt"
	"time"
)

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

func GetApprovedCommentByPostID(db *sql.DB, postID int) ([]Comment, error) {
	query := "SELECT id, post_id, author_name, author_email, content, status, created_at FROM comments WHERE post_id=? AND status='approved'"
	rows, err := db.Query(query, postID)
	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer select %w", err)
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var c Comment

		err := rows.Scan(&c.ID, &c.PostID, &c.AuthorName, &c.AuthorEmail, &c.Content, &c.Status, &c.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("erro ao scan em comments %w", err)
		}
		comments = append(comments, c)
	}
	return comments, rows.Err()
}

func GetPendingComments(db *sql.DB) ([]Comment, error) {
	query := "SELECT c.id, c.post_id, c.author_name, c.author_email, c.content, c.status, c.created_at, p.title, p.slug FROM comments c JOIN posts p ON c.post_id = p.id WHERE c.status = 'pending'"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer o select em comments %w", err)
	}
	defer rows.Close()

	var comments []Comment

	for rows.Next() {
		var c Comment
		err := rows.Scan(&c.ID, &c.PostID, &c.AuthorName, &c.AuthorEmail, &c.Content, &c.Status, &c.CreatedAt, &c.PostTitle, &c.PostSlug)
		if err != nil {
			return nil, fmt.Errorf("erro ao fazer o scan %w", err)
		}
		comments = append(comments, c)
	}
	return comments, rows.Err()
}

func CreateComment(db *sql.DB, postID int, authorName, authorEmail, content string) error {
	query := "INSERT INTO comments (post_id, author_name, author_email, content) VALUES(?, ?, ?, ?)"
	_, err := db.Exec(query, postID, authorName, authorEmail, content)
	if err != nil {
		return err
	}
	return nil
}

func ApproveComment(db *sql.DB, id int) error {
	query := "UPDATE comments SET status = 'approved' WHERE id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(db *sql.DB, id int) error {
	query := "DELETE FROM comments WHERE id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
