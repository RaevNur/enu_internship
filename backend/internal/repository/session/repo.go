package session

import (
	"database/sql"
	"fmt"

	"enu_internship/internal/models"
)

func (s *SessionRepo) Create(session *models.Session) error {
	query := `INSERT INTO sessions (uuid, user_id) 
	VALUES (?, ?, ?);`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("SessionRepo.Create: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec((*session).Uuid, (*session).UserId)
	if err != nil {
		return fmt.Errorf("SessionRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("SessionRepo.Create: %w", err)
	}

	(*session).Id = lastId
	return nil
}

func (s *SessionRepo) Delete(id int64) error {
	query := `DELETE FROM sessions WHERE id = ?`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("SessionRepo.Delete: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("SessionRepo.Delete: %w", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("SessionRepo.Delete: %w", err)
	}
	if affect != 1 {
		return fmt.Errorf("SessionRepo.Delete affected rows more than 1: %d", affect)
	}

	return nil
}

func (s *SessionRepo) GetByUserId(userId int64) (*models.Session, error) {
	query := `SELECT id, uuid, user_id, created_at FROM sessions WHERE user_id = ?`
	row := s.db.QueryRow(query, userId)

	session := models.Session{}

	err := row.Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("SessionRepo.GetByUserId: %w", err)
	}

	return &session, nil
}

func (s *SessionRepo) GetByUuid(uuid string) (*models.Session, error) {
	query := `SELECT id, uuid, user_id, created_at FROM sessions WHERE uuid = ?`
	row := s.db.QueryRow(query, uuid)

	session := models.Session{}

	err := row.Scan(&session.Id, &session.Uuid, &session.UserId, &session.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("SessionRepo.GetByUuid: %w", err)
	}

	return &session, nil
}
