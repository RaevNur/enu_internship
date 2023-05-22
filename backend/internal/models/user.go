package models

import "time"

type User struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type IUserRepo interface {
	Create(user *User) error
	GetByID(id int64) (*User, error)
	GetPassword(username string) (*User, error)
	UserExist(username string) (bool, error)
}

type IUserService interface {
	Register(user *User) error
	Login(user *User) error
	GetByID(id int64) (*User, error)
}
