package repository

import "github.com/NKV510/url_saver/internal/models"

type Repository interface {
	GetUser(firstName, lastName string) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(id int) error
	GetAllURL(userID int) ([]models.UserURL, error)
	AddURL(userID int, name, url string) error
	DeleteURL(userID int, name string) error
	GetURL(userID int, name string) error
}
