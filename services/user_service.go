package services

import (
	"github.com/travis40508/bookstore_users_api/domain/users"
	"github.com/travis40508/bookstore_users_api/utils/errors"
)

// all business logic is handled here
// you should be able to test all logic by mocking this file

// an error should always be the last thing you return, even if it returns 50 values
// we return pointers because nil refers to a memory address, so one of the two will always be nil
// either a nil user and error, or a user and nil error
// never both non-nil
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(id int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: id}

	// we call the Get() function on this user, in the user_dao
	// if it finds the user, it populates the values
	// if there's an error, it returns an error
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
