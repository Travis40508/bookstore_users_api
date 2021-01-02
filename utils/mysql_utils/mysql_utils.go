package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/travis40508/bookstore_users_api/utils/errors"
	"strings"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	// trying to cast our error to a mysql error, which also follows the error interface
	// sometimes they don't return an error of this form, though
	sqlErr, ok := err.(*mysql.MySQLError)

	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}

	return errors.NewInternalServerError("error processing request")
}
