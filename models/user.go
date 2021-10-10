package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string
	Password    string
	Email       string
	PhoneNumber string
	BirthDate   string
}

type UserDTO struct {
	ID          uint   `json:"id,omitempty"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	BirthDate   string `json:"birthDate,omitempty"`
}

func ToUser(dto UserDTO) User {
	return User{Username: dto.Username, Password: dto.Password, Email: dto.Email, BirthDate: dto.BirthDate, PhoneNumber: dto.PhoneNumber}
}

func ToUserDTO(user User) UserDTO {
	return UserDTO{ID: user.ID, Username: user.Username, Password: user.Password, Email: user.Email, BirthDate: user.BirthDate, PhoneNumber: user.PhoneNumber}
}
