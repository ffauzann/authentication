package constant

import (
	"golang.org/x/exp/slices"
)

type UserIdType string

// User identifier types
const (
	UserIdTypeUsername    UserIdType = "username"
	UserIdTypeEmail       UserIdType = "email"
	UserIdTypePhoneNumber UserIdType = "phone_number"
)

var ValidUserIdTypes = []UserIdType{UserIdTypeUsername, UserIdTypeEmail, UserIdTypePhoneNumber}

func (u *UserIdType) Validate() error {
	if !slices.Contains(ValidUserIdTypes, *u) {
		return ErrInvalidUserIdType
	}
	return nil
}
