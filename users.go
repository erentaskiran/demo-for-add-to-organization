package main

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *User) (sql.Result, error) {
	query := `INSERT INTO users (id, first_name, last_name, email, created_at, updated_at) 
              VALUES ($1, $2, $3, $4, $5, $6)`

	result, err := r.db.Exec(query, user.Id, user.FirstName, user.LastName, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to add user to database: %w", err)
	}
	return result, nil
}

func (r *UserRepository) GetUserWithEmail(email string) (User, error) {
	query := `SELECT id, first_name, last_name, email FROM users WHERE email=$1`
	row := r.db.QueryRow(query, email)

	var user User
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
	if err == sql.ErrNoRows {
		return User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return User{}, fmt.Errorf("failed to get user from database: %s", err)
	}

	return user, nil
}
