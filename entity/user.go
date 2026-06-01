package entity

import (
	vo "api-golang/value_object"
)

type User struct {
	ID    int
	Name  string
	Email *vo.Email
}

func NewUser(name string, email string) *User {
	emailVO := vo.NewEmail(email)
	return &User{
		Name:  name,
		Email: emailVO,
	}
}
