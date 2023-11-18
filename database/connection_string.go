package database

import "fmt"

// ConnectionString contains the parameters needed to make a database connection
type ConnectionString struct {
	username     string
	password     string
	host         string
	port         int
	databaseName string
}

// NewConnectionString creates a connection string ensuring you supply all the required parameters
func NewConnectionString(
	username string,
	password string,
	host string,
	port int,
	databaseName string,
) *ConnectionString {
	return &ConnectionString{
		username, password, host, port, databaseName,
	}
}

// String returns the connection string formatted correctly to make a database connection
func (cs ConnectionString) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cs.username,
		cs.password,
		cs.host,
		cs.port,
		cs.databaseName)
}

// Get returns the connection string in the correct format to make a database connection
func (cs ConnectionString) Get() string {
	return cs.String()
}
