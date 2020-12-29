package users

import (
	"fmt"
	"github.com/travis40508/bookstore_users_api/datasources/mysql/users_db"
	"github.com/travis40508/bookstore_users_api/utils/date_utils"
	"github.com/travis40508/bookstore_users_api/utils/errors"
)

// used for persisting and retrieving/accessing objects
// this is the only point where we interact with database

var (
	// retrieves a user based on id
	usersDB = make(map[int64]*User)
)

// there are advantages to making these methods, rather than functions
// return user on success, or rest error on fail
// we pass in a pointer so we can modify that object, directly
// so if an error isn't returned, we know we now have a valid user from the database
func (user *User) Get() *errors.RestErr {
	// this makes us import our users_db package
	// which call call 'init()', running the code in there
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

// this needs to be a pointer, since we're mutating the struct that we have this method on
// if it wasn't a pointer, it'd only mutate a copy
func (user *User) Save() *errors.RestErr {
	// we already have a user
	current := usersDB[user.Id]
	if current != nil {
		// email is already registered
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()

	usersDB[user.Id] = user
	return nil
}
