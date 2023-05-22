package repository

import (
	"database/sql"

	"enu_internship/internal/models"
	"enu_internship/internal/repository/session"
	"enu_internship/internal/repository/user"
)

type Repository struct {
	User    models.IUserRepo
	Session models.ISessionRepo
}

func NewRepo(db *sql.DB) *Repository {
	return &Repository{
		User:    user.NewUserRepo(db),
		Session: session.NewSessionRepo(db),
	}
}
