package repository

import (
	"fmt"
	todo "github.com/Vladosya/go-test-rest"
	"github.com/jmoiron/sqlx"
)

type PostPostgres struct {
	db *sqlx.DB
}

func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{
		db: db,
	}
}

func (r *PostPostgres) CreatePost(u todo.Post) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id",
		postsTable,
	)
	if err := r.db.QueryRow(query, u.Title, u.Description).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *PostPostgres) GetPosts() ([]todo.Post, error) {
	rows, err := r.db.Query("SELECT id, title, description FROM post")
	if err != nil {
		return []todo.Post{}, err
	}
	defer rows.Close()
	var posts []todo.Post
	for rows.Next() {
		var p todo.Post
		if err := rows.Scan(&p.Id, &p.Title, &p.Description); err != nil {
			return []todo.Post{}, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func (r *PostPostgres) DeletePostById(id string) ([]todo.Post, error) {
	rows, err := r.db.Query("DELETE FROM post WHERE id = $1 RETURNING id, title, description", id)
	if err != nil {
		return []todo.Post{}, err
	}
	defer rows.Close()
	var postDeleted []todo.Post
	for rows.Next() {
		var p todo.Post
		if err := rows.Scan(&p.Id, &p.Title, &p.Description); err != nil {
			return []todo.Post{}, err
		}
		postDeleted = append(postDeleted, p)
	}
	return postDeleted, nil
}
