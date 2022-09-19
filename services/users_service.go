package services

import (
	"github.com/mohammadshekari/bookstore_users-api/domain/users"
	"github.com/mohammadshekari/bookstore_users-api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	// Get User Original Data from DataBase
	current, err := GetUser(user.ID)
	// check User Exists
	if err != nil {
		return nil, err
	}

	// Validate user Values
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		// Set new Values to user who got from DataBase
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}
	// Send to Update Method
	if err := current.Update(); err != nil {
		return nil, err
	}
	// Return user Sent data
	return current, nil
}
