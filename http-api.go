package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var users []User = []User{}

func main() {
	router := mux.NewRouter()
	// Todo
	// access request header
	// Connect to database
	// Parse header token middleware
	// post /login
	// post /authenticate
	// post /user?pageSize=10&page=1
	// get /user?format="csv"
	// get download/user?format="csv"
	// get download/user?format="txt"
	// get download/user?format="xml"
	// patch /user/{id}
	// delete /user/{id}
	router.HandleFunc("/", test)
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{userID}", getUser).Methods("GET")
	router.HandleFunc("/users/{userID}", updateUser).Methods("PUT")
	// logger for every request
	// centralised error processing
	fmt.Println("Server started...")
	fmt.Println("Mode: ", os.Getenv("MODE"))
	fmt.Println("Location: ", ":"+os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
	// Unit tests
	// Gitlab CI
	// Deploy on heroku / firebase
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CustomResponse{"Server is active..."})
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = len(users)
	users = append(users, newUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
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

func updateUser(w http.ResponseWriter, r *http.Request) {
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
