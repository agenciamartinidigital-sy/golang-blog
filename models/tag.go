package models

import (
	"database/sql"
	"fmt"
)

func GetAllTags(db *sql.DB) ([]Tag, error) {
	query := "SELECT id, name, slug, created_at, updated_at FROM tags"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar select de tags: %w", err)
	}
	defer rows.Close()

	var tags []Tag

	for rows.Next() {
		var t Tag

		err := rows.Scan(&t.ID, &t.Name, &t.Slug, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("erro ao fazer select em tags: %w", err)
		}
		tags = append(tags, t)
	}
	return tags, rows.Err()

}

func GetTagBySlug(db *sql.DB, slug string) (Tag, error) {
	query := "SELECT id, name, slug, created_at, updated_at FROM tags WHERE slug=?"

	var s Tag
	err := db.QueryRow(query, slug).Scan(&s.ID, &s.Name, &s.Slug, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return Tag{}, err
	}
	return s, nil
}

func GetTagsByPostID(db *sql.DB, id int) ([]Tag, error) {
	query := "SELECT t.id, t.name, t.slug, t.created_at, t.updated_at FROM tags t JOIN post_tags pt ON t.id = pt.tag_id WHERE pt.post_id = ?"

	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer select em tags: %w", err)
	}
	defer rows.Close()

	var tags []Tag

	for rows.Next() {
		var t Tag

		err := rows.Scan(&t.ID, &t.Name, &t.Slug, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("esse ao fazer scan: %w", err)
		}
		tags = append(tags, t)
	}
	return tags, rows.Err()
}

func CreateTag(db *sql.DB, name, slug string) error {
	query := "INSERT INTO tags (name, slug) VALUES(?, ?)"
	_, err := db.Exec(query, name, slug)
	if err != nil {
		return err
	}
	return nil
}

func UpdateTag(db *sql.DB, id int, name, slug string) error {
	query := "UPDATE tags SET name=?, slug=? WHERE id=?"
	_, err := db.Exec(query, name, slug, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTag(db *sql.DB, id int) error {
	query := "DELETE FROM tags WHERE id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
