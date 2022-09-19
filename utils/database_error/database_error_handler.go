package database_error

import (
	errors2 "errors"
	"fmt"
	"github.com/mohammadshekari/bookstore_users-api/utils/errors"
	"gorm.io/gorm"
	"strings"
)

func NewDataBaseError(err error, model string) *errors.RestErr {
	if model == "" {
		model = "object"
	}
	if errors2.Is(err, gorm.ErrRecordNotFound) {
		return errors.NewNotFoundError(fmt.Sprintf("%s not found", model))
	}
	if errors2.Is(err, gorm.ErrRegistered) {
		return errors.NewNotFoundError(fmt.Sprintf("%s ErrRegistered", model))
	}
	if strings.Contains(err.Error(), "UNIQUE") {
		tableField := strings.Split(err.Error(), ".")
		if len(tableField) != 2 {
			return errors.NewBadRequestError(fmt.Sprintf("this value already taken"))
		}
		return errors.NewBadRequestError(fmt.Sprintf("this %s already taken", tableField[1]))
	}
	return errors.NewBadRequestError(fmt.Sprintf("unknown error happened in %s", model))
}
