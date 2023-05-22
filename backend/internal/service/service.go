package service

import (
	"enu_internship/internal/models"
	"enu_internship/internal/repository"
	"enu_internship/internal/service/session"
	"enu_internship/internal/service/user"
)

type Service struct {
	User    models.IUserService
	Session models.ISessionService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:    user.NewUserService(repo.User),
		Session: session.NewSessionService(repo.Session),
	}
}
