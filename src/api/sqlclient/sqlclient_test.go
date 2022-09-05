package sqlclient

import (
	"errors"
	"fmt"
	"testing"
)

const (
	queryGetUser = "SELECT id, email FROM users WHERE id=%d"
)

type user struct {
	Id    int64
	Email string
}

func TestSQLClient(t *testing.T) {
	user, err := getUser(123)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Email)
}

func getUser(id int64) (*user, error) {
	StartMockServer()
	testDBClient, err := Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "127.0.0.1:3306", "users"))
	if err != nil {
		panic(err)
	}

	AddMock(Mock{
		Query:   fmt.Sprintf(queryGetUser, id),
		Args:    []interface{}{1},
		Error:   nil,
		Columns: []string{"id", "email"},
		Rows: [][]interface{}{
			{1, "email1"},
			{2, "email2"},
		},
	})

	rows, err := testDBClient.Query(fmt.Sprintf(queryGetUser, id))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user user
	for rows.HasNext() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
		return &user, nil
	}

	return nil, errors.New("user not found")
}
