package user

import (
	"database/sql"
	"fmt"

	"enu_internship/internal/models"
)

func (u *UserRepo) Create(user *models.User) error {
	query := `INSERT INTO users (username, email, password, firstname, lastname) 
	VALUES (?, ?, ?, ?, ?);`

	stmt, err := u.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("UserRepo.Create: %w", err)
	}

	defer stmt.Close()

	res, err := stmt.Exec((*user).Username, (*user).Email, (*user).Password, (*user).Firstname, (*user).Lastname)
	if err != nil {
		return fmt.Errorf("UserRepo.Create: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("UserRepo.Create: %w", err)
	}

	(*user).Id = lastId
	return nil
}

func (u *UserRepo) GetByID(id int64) (*models.User, error) {
	query := `SELECT id, username, email, created_at, firstname, lastname FROM users WHERE id = ?`
	row := u.db.QueryRow(query, id)

	user := models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.Firstname, &user.Lastname)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("UserRepo.GetById: %w", err)
	}

	return &user, nil
}

func (u *UserRepo) GetPassword(nickname string) (*models.User, error) {
	query := `SELECT id, password FROM users 
	WHERE username = ?`
	row := u.db.QueryRow(query, nickname)

	user := models.User{}

	err := row.Scan(&user.Id, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("UserRepo.GetPassword: %w", err)
	}

	return &user, nil
}

func (u *UserRepo) UserExist(username string) (bool, error) {
	query := `SELECT COUNT(*) FROM users WHERE username = ?`
	row := u.db.QueryRow(query, username)

	var count int

	err := row.Scan(&count)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("UserRepo.UserExist: %w", err)
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}
