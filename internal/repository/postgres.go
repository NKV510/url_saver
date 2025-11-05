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
GetUserById(id int) (*models.User, error)

	CreateUser(user *models.User) error
	DeleteUser(id int) error
	GetAllURL(userID int) ([]models.UserURL, error)
	AddURL(userID int, name, url string) error
	DeleteURL(userID int, name string) error
*/
func (r *PsqlRepo) GetUserById(user *models.User) error {
	query := `
	INSERT INTO user (first_name, last_name)
	VALUES ($1, $2)
	RETURNING id, created_at`
	return r.db.QueryRow(context.Background(), query, user.FirstName, user.LastName).Scan(&user.ID, &user.CreatedAt)
}
func (r *PsqlRepo) DeleteUser(id int) error {
	query := `
	DELETE FROM user WHERE id=$1 AND name=$2`
	if _, err := r.db.Exec(context.Background(), query, id); err != nil {
		return fmt.Errorf("error in Postgres.go  DeleteUser: %s", err)
	}
	return nil
}
