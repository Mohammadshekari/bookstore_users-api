package users

import (
	"github.com/mohammadshekari/bookstore_users-api/utils/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	DateCreated string `json:"dateCreated"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}
