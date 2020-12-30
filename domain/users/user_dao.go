package users

import (
	"fmt"
	"github.com/travis40508/bookstore_users_api/datasources/mysql/users_db"
	"github.com/travis40508/bookstore_users_api/utils/date_utils"
	"github.com/travis40508/bookstore_users_api/utils/errors"
	"strings"
)

// used for persisting and retrieving/accessing objects
// this is the only point where we interact with database

const (
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

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
	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	// close as quickly as possible, this will be called at the end of the function block, so it's not missed
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		// we see this whenever we get our error on postman
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	// this block is the equivalent as the above, but it's not as performant as the above, nor does it have the error-handling, or reusability
	//result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	user.Id = userId

	// no need to return a user, since we're passing in a pointer reference
	return nil
}
