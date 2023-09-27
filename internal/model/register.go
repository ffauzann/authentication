package model

import "database/sql"

type RegisterRequest struct {
	Name string

	Email       string
	PhoneNumber string
	Username    string

	PlainPassword  string
	UserPassword   string
	MasterPassword string
}

type RegisterResponse struct {
	StatusCode int
	Reasons    []string
}

func (r *RegisterRequest) ToUser() *User {
	return &User{
		Name:  r.Name,
		Email: r.Email,
		PhoneNumber: sql.NullString{
			String: r.PhoneNumber,
			Valid:  r.PhoneNumber != "",
		},
		Username: sql.NullString{
			String: r.Username,
			Valid:  r.Username != "",
		},

		Password:       r.UserPassword,
		MasterPassword: r.MasterPassword,
	}
}
