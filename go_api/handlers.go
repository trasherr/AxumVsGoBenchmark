package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userRequest UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	email := userRequest.Email
	password := userRequest.Password

	if email == "" || password == "" {
		http.Error(w, "Missing email or password", http.StatusBadRequest)
		return
	}

	user, err := userModel.FindByEmailPassword(email, password)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Here, you should check the password. This example assumes a plaintext check,
	// but in a real application, you should hash and compare the stored hash.
	if user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	setJsonResponseHeader(w)
	json.NewEncoder(w).Encode(user)

}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	print(r.Body)
	var userRequest UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Create a new user object
	user := &User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Age:      userRequest.Age,
		Password: userRequest.Password, // Store password securely in a real application
	}

	// Use the UserModel to create a new user
	err = userModel.Create(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	setJsonResponseHeader(w)
	json.NewEncoder(w).Encode(user)

}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users, err := userModel.GetAll()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	setJsonResponseHeader(w)
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userRequest UserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user := &User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Age:      userRequest.Age,
		Password: userRequest.Password,
	}

	user, errr := userModel.Update(user)
	if errr != nil {
		log.Println(errr.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	setJsonResponseHeader(w)
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var userloginRequest UserLoginRequest
	err := json.NewDecoder(r.Body).Decode(&userloginRequest)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	user, err := userModel.Delete(userloginRequest.Email, userloginRequest.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	setJsonResponseHeader(w)
	json.NewEncoder(w).Encode(user)
}
