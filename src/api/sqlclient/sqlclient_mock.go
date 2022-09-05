package sqlclient

import "errors"

func AddMock(mock Mock) {
	if dbClient == nil {
		return
	}

	client, ok := dbClient.(*clientMock)
	if !ok {
		return
	}

	if client.mocks == nil {
		client.mocks = make(map[string]Mock, 0)
	}
	client.mocks[mock.Query] = mock
}

type clientMock struct {
	mocks map[string]Mock
}

type Mock struct {
	Query   string
	Args    []interface{}
	Error   error
	Columns []string
	Rows    [][]interface{}
}

func (client *clientMock) Query(query string, args ...interface{}) (rows, error) {
	mock, exist := client.mocks[query]
	if !exist {
		return nil, errors.New("no mock found")
	}

	if mock.Error != nil {
		return nil, mock.Error
	}

	rows := rowMock{
		Columns: mock.Columns,
		Rows:    mock.Rows,
	}

	return &rows, nil
}
