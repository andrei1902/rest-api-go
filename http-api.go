package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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
	router.HandleFunc("/", Test)
	router.HandleFunc("/users", AddUser).Methods("POST")
	router.HandleFunc("/users", GetUsers).Methods("GET")
	router.HandleFunc("/users/{userID}", GetUser).Methods("GET")
	router.HandleFunc("/users/{userID}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{userID}", PatchUser).Methods("PATCH")
	router.HandleFunc("/users/{userID}", DeleteUser).Methods("DELETE")
	loggingRouter := LoggerMiddleware(router)
	// centralised error processing
	fmt.Println("Server started...")
	fmt.Println("Mode: ", os.Getenv("MODE"))
	fmt.Println("Location: ", ":"+os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), loggingRouter))
	// Unit tests
	// Gitlab CI
	// Deploy on heroku / firebase
}
