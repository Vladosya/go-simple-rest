package repository

import (
	"fmt"
	todo "github.com/Vladosya/go-test-rest"
	"github.com/jmoiron/sqlx"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{
		db: db,
	}
}

func (r *UserPostgres) CreateUser(u todo.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (name, surname) VALUES ($1, $2) RETURNING id",
		usersTable,
	)
	if err := r.db.QueryRow(query, u.Name, u.Surname).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *UserPostgres) GetUsers() ([]todo.User, error) {
	rows, err := r.db.Query("SELECT * FROM person")
	if err != nil {
		return []todo.User{}, err
	}
	defer rows.Close()
	var users []todo.User
	for rows.Next() {
		var p todo.User
		if err := rows.Scan(&p.Id, &p.Name, &p.Surname); err != nil {
			return []todo.User{}, err
		}
		users = append(users, p)
	}
	return users, nil
}

func (r *UserPostgres) GetUserById(id string) ([]todo.User, error) {
	rows, err := r.db.Query("SELECT * FROM person WHERE id = $1", id)
	if err != nil {
		return []todo.User{}, err
	}
	defer rows.Close()
	var user []todo.User
	for rows.Next() {
		var p todo.User
		if err := rows.Scan(&p.Id, &p.Name, &p.Surname); err != nil {
			return []todo.User{}, err
		}
		user = append(user, p)
	}

	return user, nil
}

func (r *UserPostgres) DeleteUserById(id string) ([]todo.User, error) {
	rows, err := r.db.Query("DELETE FROM person WHERE id = $1 RETURNING *", id)
	if err != nil {
		return []todo.User{}, err
	}
	defer rows.Close()
	var userDeleted []todo.User
	for rows.Next() {
		var p todo.User
		if err := rows.Scan(&p.Id, &p.Name, &p.Surname); err != nil {
			return []todo.User{}, err
		}
		userDeleted = append(userDeleted, p)
	}
	return userDeleted, nil
}
