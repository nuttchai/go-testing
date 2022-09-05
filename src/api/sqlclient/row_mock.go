package sqlclient

import "errors"

type rowMock struct {
	Columns []string
	Rows    [][]interface{}

	currentIndex int
}

type sqlRowsMock struct {
	rows rowMock
}

func (r *rowMock) HasNext() bool {
	return r.currentIndex < len(r.Rows)
}

func (r *rowMock) Close() error {
	return nil
}

func (r *rowMock) Scan(destinations ...interface{}) error {
	mockedRow := r.Rows[r.currentIndex]
	if len(mockedRow) != len(destinations) {
		return errors.New("Invalid destinations len")
	}

	for index, value := range mockedRow {
		destinations[index] = value
	}

	return nil
}
