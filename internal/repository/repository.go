package repository

import "github.com/NKV510/url_saver/internal/models"

type Repository interface {
	GetUserById(id int) (*models.User, error)
	CreateUser(user *models.User) error
	DeleteUser(id int) error
	GetAllURL(userID int) ([]models.UserURL, error)
	AddURL(userID int, name, url string) error
	DeleteURL(userID int, name string) error
}
