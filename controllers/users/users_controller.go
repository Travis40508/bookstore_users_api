package users

import (
	"github.com/gin-gonic/gin"
	"github.com/travis40508/bookstore_users_api/domain/users"
	"github.com/travis40508/bookstore_users_api/services"
	"github.com/travis40508/bookstore_users_api/utils/errors"
	"net/http"
	"strconv"
)

// entry point of every request
// take request, validate if we have all of params, then send the handling to the service
// so all business logic is in services
// each request will be handled on a different goroutine. for this reason we want exclusively stateless controllers
// if we had member variables here where we depended on logic, they'd all be modified on different threads

// handles all requests that attempt to create a user
func CreateUser(c *gin.Context) {
	var user users.User
	// this guy is doing everything the below, commented-out code is doing
	if err := c.ShouldBindJSON(&user); err != nil {
		// this will throw an error if the user interface doesn't match up to the types passed into the request body
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// handles all requests that attempt to get a user from the database
// you should always check for a bad request first
func GetUser(c *gin.Context) {
	// user_id comes from the urlMappings() param value
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, user)
}

//func SearchUser(c *gin.Context) {
//	c.String(http.StatusNotImplemented, "implementation missing")
//}

// THIS IS THE SAME THING AS c.shouldBindJson(&user)
// getting byte array from request, this is the JSON body
// this request comes from the request interacting with our endpoint in our mapUrls() function
//bytes, err := ioutil.ReadAll(c.Request.Body)
//
//if err != nil {
//	return
//}
//
//// UnMarshal takes a byte-array and pointer to an interface, so we have to refer to the memory address
//// it maps the values from the request body onto our struct object
//if err := json.Unmarshal(bytes, &user); err != nil {
//	return
//}
