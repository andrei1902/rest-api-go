package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var data []User = []User{}

func main() {
	router := mux.NewRouter()
	// Todo
	// Set server status from env variables
	// access request payload
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
	// get /user/{id}
	// put /user/{id}
	// patch /user/{id}
	// put /user/{id}
	// delete /user/{id}
	router.HandleFunc("/", test)
	router.HandleFunc("/users", addUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	// logger for every request
	// centralised error processing
	log.Fatal(http.ListenAndServe(":4500", router))
	fmt.Println("Server started...")
	// Unit tests
	// Gitlab CI
	// Deploy on heroku / firebase
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CustomResponse{"Server is active..."})
}

func addUser(w http.ResponseWriter, r *http.Request) {
	// requestParam := mux.Vars(r)["item"]
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = len(data) + 1
	data = append(data, newUser)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
