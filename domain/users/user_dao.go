package users

import (
	"github.com/travis40508/bookstore_users_api/datasources/mysql/users_db"
	"github.com/travis40508/bookstore_users_api/utils/date_utils"
	"github.com/travis40508/bookstore_users_api/utils/errors"
	"github.com/travis40508/bookstore_users_api/utils/mysql_utils"
)

// used for persisting and retrieving/accessing objects
// this is the only point where we interact with database

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

// there are advantages to making these methods, rather than functions
// return user on success, or rest error on fail
// we pass in a pointer so we can modify that object, directly
// so if an error isn't returned, we know we now have a valid user from the database
// Get always looks based on primary key/id, 'Find' always looks based on parameters
func (user *User) Get() *errors.RestErr {
	// this makes us import our users_db package
	// which call call 'init()', running the code in there
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	// this is the parameter we pass in so that our query statement above can work
	// it should be noted that if you query more than one user (with Query(), rather than QueryRow()), that it once again
	// opens the DB and must be closed on defer
	result := stmt.QueryRow(user.Id)
	// we're sending a pointer because if we don't pass a pointer, the scan will just assign these values
	// to a copy, rather than to the user we're adding this method on
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		// this means we have an error when trying to retrieve this row from the database
		// the scan takes our result from our query
		return mysql_utils.ParseError(err)
	}

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

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		// this attempts to convert saveErr into a MySqlErr, since both implement the Error interface
		// they don't always return a mysql error, though. sometimes they'll just return a string, it's only an attempt
		return mysql_utils.ParseError(saveErr)
	}

	// this block is the equivalent as the above, but it's not as performant as the above, nor does it have the error-handling, or reusability
	//result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	user.Id = userId

	// no need to return a user, since we're passing in a pointer reference
	return nil
}
