package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// AddUser - controller for POST /users
func AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = len(users)
	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// GetUsers - controller for GET /users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUser - controller for GET /users/ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["userID"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("userID should be an integer, dumbass!"))
		return
	}
	if id >= len(users) {
		w.WriteHeader(404)
		w.Write([]byte("User, not found, my killa!"))
		return
	}
	user := users[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser - controller for POST /users/ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["userID"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("userID should be an integer, dumbass!"))
		return
	}
	if id >= len(users) {
		w.WriteHeader(404)
		w.Write([]byte("User, not found, my killa!"))
		return
	}

	var updatedU User
	json.NewDecoder(r.Body).Decode(&updatedU)

	users[id] = updatedU

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedU)
}

// PatchUser - controller for PATCH /users/ID
func PatchUser(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["userID"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("userID should be an integer, dumbass!"))
		return
	}
	if id >= len(users) {
		w.WriteHeader(404)
		w.Write([]byte("User, not found, my killa!"))
		return
	}

	user := users[id]
	json.NewDecoder(r.Body).Decode(&user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// DeleteUser - controller for Delete /users/ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["userID"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("userID should be an integer, dumbass!"))
		return
	}
	if id >= len(users) {
		w.WriteHeader(404)
		w.Write([]byte("User, not found, my killa!"))
		return
	}

	users = append(users[:id], users[id+1:]...)

	w.Write([]byte("The user was deleted"))
}
