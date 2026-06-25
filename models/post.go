package models

import (
	"database/sql"
	"fmt"
	"time"
)

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

func GetPublishedPosts(db *sql.DB) ([]Post, error) {
	query := "SELECT p.id, p.category_id, p.title, p.slug, p.content, p.status, p.published_at, p.created_at, p.updated_at, c.id, c.name, c.slug FROM posts p JOIN categories c ON p.category_id = c.id WHERE p.status ='published' ORDER BY p.published_at DESC"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer select em post: %w", err)
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.CategoryID, &p.Title, &p.Slug, &p.Content, &p.Status, &p.PublishedAt, &p.CreatedAt, &p.UpdatedAt, &p.Category.ID, &p.Category.Name, &p.Category.Slug)
		if err != nil {
			return nil, fmt.Errorf("erro ao scan %w", err)
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func GetAllPostsAdmin(db *sql.DB) ([]Post, error) {
	query := "SELECT p.id, p.category_id, p.title, p.slug, p.content, p.status, p.published_at, p.created_at, p.updated_at, c.id, c.name, c.slug FROM posts p JOIN categories c ON p.category_id = c.id ORDER BY p.published_at DESC"

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Erro ao fazer o select: %w", err)
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.CategoryID, &p.Title, &p.Slug, &p.Content, &p.Status, &p.PublishedAt, &p.CreatedAt, &p.UpdatedAt, &p.Category.ID, &p.Category.Name, &p.Category.Slug)
		if err != nil {
			return nil, fmt.Errorf("erro ao fazer o scan: %w", err)
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func GetPostsByCategory(db *sql.DB, slug string) ([]Post, error) {
	query := "SELECT p.id, p.category_id, p.title, p.slug, p.content, p.status, p.published_at, p.created_at, p.updated_at, c.id, c.name, c.slug FROM posts p JOIN categories c ON p.category_id = c.id WHERE c.slug=? AND p.status ='published' ORDER BY p.published_at DESC"

	rows, err := db.Query(query, slug)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer select: %w", err)
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.CategoryID, &p.Title, &p.Slug, &p.Content, &p.Status, &p.PublishedAt, &p.CreatedAt, &p.UpdatedAt, &p.Category.ID, &p.Category.Name, &p.Category.Slug)
		if err != nil {
			return nil, fmt.Errorf("Erro ao fazer o scan %w", err)
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func GetPostsByTag(db *sql.DB, slug string) ([]Post, error) {
	query := "SELECT p.id, p.category_id, p.title, p.slug, p.content, p.status, p.published_at, p.created_at, p.updated_at, c.id, c.name, c.slug FROM posts p JOIN categories c ON p.category_id = c.id JOIN post_tags pt ON p.id = pt.post_id JOIN tags t ON pt.tag_id = t.id WHERE t.slug=? AND p.status ='published'"

	rows, err := db.Query(query, slug)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer o select: %w", err)
	}
	defer rows.Close()

	var posts []Post

	for rows.Next() {
		var p Post
		err := rows.Scan(&p.ID, &p.CategoryID, &p.Title, &p.Slug, &p.Content, &p.Status, &p.PublishedAt, &p.CreatedAt, &p.UpdatedAt, &p.Category.ID, &p.Category.Name, &p.Category.Slug)
		if err != nil {
			return nil, fmt.Errorf("erro ao fazer o scan: %w", err)
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func GetPostBySlug(db *sql.DB, slug string) (Post, error) {
	query := "SELECT p.id, p.category_id, p.title, p.slug, p.content, p.status, p.published_at, p.created_at, p.updated_at, c.id, c.name, c.slug FROM posts p JOIN categories c ON p.category_id = c.id WHERE p.slug = ?"

	var p Post
	err := db.QueryRow(query, slug).Scan(&p.ID, &p.CategoryID, &p.Title, &p.Slug, &p.Content, &p.Status, &p.PublishedAt, &p.CreatedAt, &p.UpdatedAt, &p.Category.ID, &p.Category.Name, &p.Category.Slug)
	if err != nil {
		return Post{}, fmt.Errorf("Erro ao fazer o scan: %w", err)
	}
	tags, err := GetTagsByPostID(db, p.ID)
	if err != nil {
		return Post{}, err
	}
	p.Tags = tags
	return p, nil
}

func CreatePost(db *sql.DB, title, content, slug, status string, categoryID int, publishedAt sql.NullTime, tagIDs []int) error {
	query := "INSERT INTO posts(title, content, slug, status, category_id, published_at) values(?, ?, ?, ?, ?, ?)"

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.Exec(query, title, content, slug, status, categoryID, publishedAt)
	if err != nil {
		return err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	for _, tagID := range tagIDs {
		_, err = tx.Exec("INSERT INTO post_tags (post_id, tag_id) VALUES (?, ?)", postID, tagID)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}

func UpdatePost(db *sql.DB, title, content, slug, status string, id, categoryID int, publishedAt sql.NullTime, tagIDs []int) error {

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("erro ao iniciar a transição: %w", err)
	}
	defer tx.Rollback()

	// Update
	updateQuery := `UPDATE posts SET title=?, content=?, slug=?, status=?, category_id=?, published_at=? WHERE id=?`

	_, err = tx.Exec(updateQuery, title, content, slug, status, categoryID, publishedAt, id)
	if err != nil {
		return fmt.Errorf("erro ao atualizar: %w", err)
	}

	// Delete
	_, err = tx.Exec("DELETE FROM post_tags WHERE post_id=?", id)
	if err != nil {
		return fmt.Errorf("erro ao deletar: %w", err)
	}

	// Insert
	insertTagQuery := "INSERT INTO post_tags(post_id, tag_id) VALUES(?, ?)"
	for _, tagID := range tagIDs {
		_, err := tx.Exec(insertTagQuery, id, tagID)
		if err != nil {
			return fmt.Errorf("erro ao inserir tag %d: %w", tagID, err)
		}
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("erro ao comitar a tansacção: %w", err)
	}

	return nil

}

func DeletePost(db *sql.DB, postID int) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("erro na transação: %w", err)
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM post_tags WHERE post_id=?", postID)
	if err != nil {
		return fmt.Errorf("erro ao fazer o delete da tag: %w", err)
	}

	_, err = tx.Exec("DELETE FROM comments WHERE post_id=?", postID)
	if err != nil {
		return fmt.Errorf("erro ao fazer o delete do comment: %w", err)
	}

	_, err = tx.Exec("DELETE FROM posts WHERE id=?", postID)
	if err != nil {
		return fmt.Errorf("erro ao fazer o delete do post: %w", err)
	}

	if err := tx.Commit(); err != nil{
		return fmt.Errorf("Erro ao fazer o commit")
	}
	return nil
}
