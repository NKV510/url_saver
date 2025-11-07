package repository

import (
	"context"
	"fmt"

	"github.com/NKV510/url_saver/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PsqlRepo struct {
	db *pgxpool.Pool
}

func NewPsqlRepo(db *pgxpool.Pool) *PsqlRepo {
	return &PsqlRepo{db: db}
}

/*
##GetUserById(id int) (*models.User, error)
##CreateUser(user *models.User) error
##DeleteUser(id int) error
##GetAllURL(userID int) ([]models.UserURL, error)
##AddURL(userID int, name, url string) error
##DeleteURL(userID int, name string) error
*/
func (r *PsqlRepo) CreateUser(user *models.User) error {
	query := `INSERT INTO user (first_name, last_name, created_at)
	VALUES ($1, $2, $3)`
	if _, err := r.db.Exec(context.Background(), query, user.FirstName, user.LastName, user.CreatedAt); err != nil {
		return fmt.Errorf("error in Postgres.go  CreateUser: %s", err)
	}
	return nil
}

func (r *PsqlRepo) DeleteUser(ID int) error {
	query := `DELETE FROM user WHERE id=$1`
	if _, err := r.db.Exec(context.Background(), query, ID); err != nil {
		return fmt.Errorf("error in Postgres.go  DeleteUser: %s", err)
	}
	return nil
}

func (r *PsqlRepo) GetUser(firstName, lastName string) (*models.User, error) {
	query := `
	SELECT id, first_name, last_name, created_at FROM users WHERE first_name = $1 AND last_name = $2`
	var user models.User
	if err := r.db.QueryRow(context.Background(), query, firstName, lastName).Scan(&user.ID, &user.FirstName, &user.LastName, &user.CreatedAt); err != nil {
		return nil, fmt.Errorf("error in Postgres.go  GetUser: %s", err)
	}
	return &user, nil
}

func (r *PsqlRepo) AddURL(userID int, name, url string) error {
	query := `INSERT INTO user_url (userID, name, url)
	VALUES ($1, $2, $3)`
	if _, err := r.db.Exec(context.Background(), query, userID, name, url); err != nil {
		return fmt.Errorf("error in Postgres.go  AddURL: %s", err)
	}
	return nil
}

func (r *PsqlRepo) GetAllURL(userID int) ([]models.UserURL, error) {
	query := `
	SELECT name, url FROM user_url WHERE user_id = $1`
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("error in conect by GetAllURL: %s", err)
	}
	defer rows.Close()
	var urls []models.UserURL
	for rows.Next() {
		var url models.UserURL
		if err := rows.Scan(&url.Name, &url.URL); err != nil {
			fmt.Println("Can not read user url", err)
		}
		urls = append(urls, url)
	}
	return urls, nil
}

func (r *PsqlRepo) GetURL(userID int, name string) (models.UserURL, error) {
	query := `
	SELECT url FROM user_url WHERE user_id = $1 AND name = $2`
	var url models.UserURL
	if err := r.db.QueryRow(context.Background(), query, userID, name).Scan(&url.URL); err != nil {
		return models.UserURL{}, fmt.Errorf("error from ger URL: %s", err)
	}
	return url, nil
}

func (r *PsqlRepo) DeleteURL(userID int, name string) error {
	query := `DELETE FROM user_url WHERE id=$1 AND name = $2`
	if _, err := r.db.Exec(context.Background(), query, userID, name); err != nil {
		return fmt.Errorf("error in Postgres.go  DeleteUser: %s", err)
	}
	return nil
}
