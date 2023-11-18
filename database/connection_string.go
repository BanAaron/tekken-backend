package database

import "fmt"

type ConnectionString struct {
	username     string
	password     string
	host         string
	port         int
	databaseName string
}

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

func (cs ConnectionString) String() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s",
		cs.username,
		cs.password,
		cs.host,
		cs.port,
		cs.databaseName)
}

func (cs ConnectionString) Get() string {
	return cs.String()
}
