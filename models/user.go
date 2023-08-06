package models

import (
	gormdb "project2/gormDB"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	Name     string `json:"name" gorm:"not null"`
	Surname  string `json:"surname" gorm:"not null"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) FindUserById(id int) (User, error) {
	var user User

	if err := gormdb.DB.Where("ID=&", id).Find(&user).Error; err != nil {
		return User{}, nil
	}

	return user, nil
}

func (u *User) CreateUser() (*User, error) {
	if err := gormdb.DB.Create(&u).Error; err != nil {
		return &User{}, err
	}
	return u, nil
}
