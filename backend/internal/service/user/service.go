package user

import (
	"fmt"

	"enu_internship/internal/helper"
	"enu_internship/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// validate email, names, password
// check if username exists
// password -> hash(password)
// createdAt -> time.Now()
func (s *UserService) Register(user *models.User) error {
	if !helper.ValidateEmail(user.Email) {
		return fmt.Errorf("UserService.Register: %w", &helper.ValidateError{
			Field:       "email",
			Description: "email must be correct",
		})
	}

	if !helper.ValidateNames(user.Firstname) {
		return fmt.Errorf("UserService.Register: %w", &helper.ValidateError{
			Field:       "Firstname",
			Description: "name cant have spaces, be empty and longer than 40 symbols",
		})
	}

	if !helper.ValidateNames(user.Lastname) {
		return fmt.Errorf("UserService.Register: %w", &helper.ValidateError{
			Field:       "Lastname",
			Description: "name cant have spaces, be empty and longer than 40 symbols",
		})
	}

	if !helper.ValidatePassword(user.Password) {
		return fmt.Errorf("UserService.Register: %w", &helper.ValidateError{
			Field:       "password",
			Description: "password cant have spaces, shorter than 7 and longer than 40 symbols",
		})
	}

	exist, err := s.repo.UserExist(user.Username)
	if err != nil {
		return fmt.Errorf("UserService.Register: %w", err)
	}

	if exist {
		return fmt.Errorf("UserService.Register: %w", &helper.ExistsError{
			Title:       "user is exists",
			Description: "user with this email or nickname already exists",
		})
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("UserService.Register: %w", err)
	}

	user.Password = string(hashed)
	return s.repo.Create(user)
}

func (s *UserService) Login(user *models.User) error {
	exist, err := s.repo.UserExist(user.Username)
	if err != nil {
		return fmt.Errorf("UserService.Login: %w", err)
	}

	if exist {
		userInfo, err := s.repo.GetPassword(user.Username)
		if err != nil {
			return fmt.Errorf("UserService.Login: %w", err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(user.Password))
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return fmt.Errorf("UserService.Login: %w", &helper.ValidateError{
				Field:       "password",
				Description: "password is not correct",
			})
		} else if err != nil {
			return fmt.Errorf("UserService.Login: %w", err)
		}

		user.Id = userInfo.Id
	} else {
		return fmt.Errorf("UserService.Login: %w", &helper.ExistsError{
			Title:       "user is not exists",
			Description: "user with this nickname in not exists",
		})
	}

	return nil
}

func (s *UserService) GetByID(id int64) (*models.User, error) {
	userInfo, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("UserService.GetByID: %w", err)
	}

	return userInfo, nil
}
