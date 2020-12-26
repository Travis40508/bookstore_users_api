package users

import (
	"github.com/travis40508/bookstore_users_api/utils/errors"
	"strings"
)

// used for moving objects between layers
// notice the tics, instead of quotation marks around json
// this tells go which json keys map to our values in our struct
// they also allow us to easily pass our struct as a json object as a response
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

// making this a method (rather than a function) so the user can validate itself
func (user *User) Validate() *errors.RestErr {
	// trim spaces of email
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	return nil
}
