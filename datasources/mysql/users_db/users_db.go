package users_db

// database/sql is part of the Go standard library
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // we're not using this import, but we're importing it so the code can understand it
	"log"
	"os"
)

const (
	mysql_users_db_username = "mysql_users_db_username"
	mysql_users_db_password = "mysql_users_db_password"
	mysql_users_db_host     = "mysql_users_db_host"
	mysql_users_db_schema   = "mysql_users_db_schema"
)

var (
	Client *sql.DB
	// these can all be found in the build config Environment variables in this IDE (this is like storing a variable in your .bash_profile)
	username = os.Getenv(mysql_users_db_username)
	password = os.Getenv(mysql_users_db_password)
	host     = os.Getenv(mysql_users_db_host)
	schema   = os.Getenv(mysql_users_db_schema)
)

// init is called by default in Go the first time this package (not file) is imported
func init() {
	// username, password, host, schema name
	// this allows us to connect to our mysql db
	// may need to make this 127.0.0.1:3006
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	// we go ahead and make Client and err early, since we only want to create it one time, and to just use it
	// everywhere in this file
	var err error
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		// this will prevent the app from starting up (as we need this db to make anything in the server to work)
		panic(err)
	}

	// pings our db
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database successfully configured")
}
