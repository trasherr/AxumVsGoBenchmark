package main

import (
	"database/sql"
	"fmt"
)
type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

// UserModel handles database operations for the User model
type UserModel struct {
	DB *sql.DB
}

// Create inserts a new user into the database
func (m *UserModel) Create(user *User) error {
	query := "INSERT INTO users(name, email, age, password) VALUES(?, ?, ?, ?)"
	result, err := m.DB.Exec(query, user.Name, user.Email, user.Age, user.Password)
	if err != nil {
		return fmt.Errorf("unable to insert user: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("unable to get last insert id: %w", err)
	}
	user.ID = int(id)
	return nil
}

// GetAll retrieves all users from the database
func (m *UserModel) GetAll() ([]User, error) {
	rows, err := m.DB.Query("SELECT id, name, email, age, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// FindByEmailPassword finds a user by their email and password
func (m *UserModel) FindByEmailPassword(email string, password string) (*User, error) {
	user := &User{}
	query := "SELECT id, name, email, age, password FROM users WHERE email = ? AND password = ?"
	err := m.DB.QueryRow(query, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}
	return user, nil
}

// Update updates the user information in the database
func (m *UserModel) Update(user *User) (*User, error) {
	query := "UPDATE users SET name = ?, password = ?, age = ? WHERE email = ?"
	_, err := m.DB.Exec(query, user.Name, user.Password, user.Age, user.Email)
	if err != nil {
		return nil, fmt.Errorf("unable to update user: %w", err)
	}

	return user, nil
}

// Delete removes a user from the database
func (m *UserModel) Delete(email string, password string) (*User, error) {
	user := &User{}
	queryFind := "SELECT id, name, email, age, password FROM users WHERE email = ? AND password = ?"
	err := m.DB.QueryRow(queryFind, email, password).Scan(&user.ID, &user.Name, &user.Email, &user.Age, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // User not found
		}
		return nil, err
	}

	query := "DELETE FROM users WHERE email = ? AND password = ?"
	result, err := m.DB.Exec(query, email, password)
	if err != nil {
		return nil, fmt.Errorf("unable to delete user: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, sql.ErrNoRows // No rows affected, user not found
	}

	return user, nil
}