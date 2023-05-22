package session

import (
	"fmt"
	"time"

	"enu_internship/internal/models"

	uuid "github.com/satori/go.uuid"
)

func (s *SessionService) GenerateSession(userId int64) (*models.Session, error) {
	u1 := uuid.NewV4()
	session := models.Session{
		Uuid:      u1.String(),
		UserId:    userId,
		CreatedAt: time.Now(),
	}

	err := s.repo.Create(&session)
	if err != nil {
		return nil, fmt.Errorf("SessionService.GenerateCookie: %w", err)
	}

	return &session, nil
}

func (s *SessionService) DeleteSession(id int64) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("SessionService.DeleteSession: %w", err)
	}

	return nil
}

func (s *SessionService) GetByUserId(userId int64) (*models.Session, error) {
	session, err := s.repo.GetByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("SessionService.GetByUserId: %w", err)
	}

	return session, nil
}

func (s *SessionService) GetByUuid(uuid string) (*models.Session, error) {
	session, err := s.repo.GetByUuid(uuid)
	if err != nil {
		return nil, fmt.Errorf("SessionService.GetByUuid: %w", err)
	}

	return session, nil
}
