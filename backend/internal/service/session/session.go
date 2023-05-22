package session

import "enu_internship/internal/models"

type SessionService struct {
	repo models.ISessionRepo
}

func NewSessionService(repo models.ISessionRepo) *SessionService {
	return &SessionService{repo}
}
