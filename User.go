package main

// User Data type
type User struct {
	ID           int
	PersonalData UserPersonalData `json:"personalData"`
	IsDeleted    bool
	PasswordHash string
}
