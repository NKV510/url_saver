package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserURL struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	URL    string `json:"url"`
}

type CreateUser struct {
	FirstName string `json:"fist_name"`
	LastName  string `json:"last_name"`
}

type CrateURL struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
