package models

import "gorm.io/gorm"

type Model interface {
	Validate() bool
}

type Contact struct {
	gorm.Model
	Email    string `json:"email" valid:"required,email"`
	Name     string `json:"name" valid:"required,alpha"`
	LastName string `json:"last_name" valid:"required,alpha"`
	Company  string `json:"company" valid:"optional"`
	Phone    string `json:"phone" valid:"numeric,optional"`
}

func(c *Contact) Validate() bool {
	return true
}

type Message struct {
	Message string `json:"msg"`
}
