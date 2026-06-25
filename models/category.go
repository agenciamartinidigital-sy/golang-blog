package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Category struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Função de leitura de banco de dados
func GetAllCategories(db *sql.DB) ([]Category, error) {
	query := "SELECT id, name, slug, created_at, updated_at FROM categories"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar select de categories: %w", err)
	}
	defer rows.Close()

	var categories []Category

	for rows.Next() {
		var c Category

		err := rows.Scan(&c.ID, &c.Name, &c.Slug, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("erro ao scannear linha da category %w", err)
		}
		categories = append(categories, c)
	}

	return categories, rows.Err()
}

func GetCategoryBySlug(db *sql.DB, slug string) (Category, error) {
	query := "SELECT id, name, slug, created_at, updated_at FROM categories WHERE slug = ?"

	var c Category
	err := db.QueryRow(query, slug).Scan(&c.ID, &c.Name, &c.Slug, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return Category{}, err
	}
	return c, nil
}

func GetCategoryByID(db *sql.DB, id int) (Category, error) {
	query := "SELECT id, name, slug, created_at, updated_at FROM categories WHERE id = ?"

	var c Category
	err := db.QueryRow(query, id).Scan(&c.ID, &c.Name, &c.Slug, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return Category{}, err
	}
	return c, nil
}

func CreateCategory(db *sql.DB, name, slug string) error {
	query := "INSERT INTO categories (name, slug) VALUES(?, ?)"

	_, err := db.Exec(query, name, slug)
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(db *sql.DB, id int, name, slug string) error {
	query := "UPDATE categories SET name=?, slug=? WHERE id=?"
	_, err := db.Exec(query, name, slug, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(db *sql.DB, id int) error {
	query := "DELETE FROM categories WHERE id=?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
