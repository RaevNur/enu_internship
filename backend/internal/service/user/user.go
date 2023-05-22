package user

import "enu_internship/internal/models"

type UserService struct {
	repo models.IUserRepo
}

func NewUserService(repo models.IUserRepo) *UserService {
	return &UserService{repo}
}
