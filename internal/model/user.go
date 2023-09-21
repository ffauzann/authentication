package model

import "database/sql"

type User struct {
	Id          uint64         `db:"id"`
	Name        string         `db:"name"`
	Email       string         `db:"email"`
	Username    sql.NullString `db:"username"`
	PhoneNumber sql.NullString `db:"phone_number"`

	Password       string `db:"password"`
	MasterPassword string `db:"master_password"`

	IsBlocked bool `db:"is_blocked"`
	Roles     []string

	Timestamp
}

type IsUserExistRequest struct {
	Username    string
	Email       string
	PhoneNumber string
}

type IsUserExistResponse struct {
	IsExist bool
	Reasons []string
}
