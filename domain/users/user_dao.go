package users

import (
	"fmt"
	"github.com/mohammadshekari/bookstore_users-api/database"
	"github.com/mohammadshekari/bookstore_users-api/utils/database_error"
	"github.com/mohammadshekari/bookstore_users-api/utils/errors"
)

func (user *User) Get() *errors.RestErr {
	if err := database.Db.First(&user, user.ID); err.Error != nil {
		return database_error.NewDataBaseError(err.Error, "user")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	if result := database.Db.First(&User{}, user.ID); result.Error == nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.ID))
	}
	//if result := database.Db.First(&User{}, "email = ?", user.Email); result.Error == nil {
	//	return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	//}
	if err := database.Db.Create(&user).Error; err != nil {
		return database_error.NewDataBaseError(err, "user")
	}
	return nil
}

func (user *User) Update() *errors.RestErr {
	// Already Checked User exists and Values of user has been Updated
	// Send to Update
	if result := database.Db.Model(&User{ID: user.ID}).Select("first_name", "last_name", "email").Updates(user); result.Error != nil {
		return errors.NewBadRequestError(fmt.Sprintf("some %s error happened", user.Email))
	}
	return nil
}
