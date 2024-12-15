package models

type User struct {
	Name   string `json:"name" validate:"required"`
	PAN    string `json:"pan" validate:"required,pan"`
	Mobile string `json:"mobile" validate:"required,mobile"`
	Email  string `json:"email" validate:"required,email"`
}
