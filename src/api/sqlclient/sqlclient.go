package sqlclient

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	goEnvironment = "GO_ENVIRONMENT"
	production    = "production"
)

var (
	isMocked bool
	dbClient SqlClient
)

type client struct {
	db *sql.DB
}

type SqlClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

func StartMockServer() {
	isMocked = true
}

func StopMockServer() {
	isMocked = false
}

func Open(driverName, dataSourceName string) (SqlClient, error) {
	if isMocked && !isProduction() {
		dbClient = &clientMock{}
		return dbClient, nil
	}

	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	dbClient = &client{
		db: db,
	}

	return dbClient, nil
}

func isProduction() bool {
	return os.Getenv(goEnvironment) == production
}

func (client *client) Query(query string, args ...interface{}) (rows, error) {
	returnedRows, err := client.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	result := sqlRows{
		rows: returnedRows,
	}

	return &result, nil
}
