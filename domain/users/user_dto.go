package users

import (
	"github.com/mohammadshekari/bookstore_users-api/utils/errors"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	ID          int64          `gorm:"primaryKey" json:"id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Email       string         `json:"email"`
	DateCreated time.Time      `gorm:"<-:create" json:"date_created"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}
