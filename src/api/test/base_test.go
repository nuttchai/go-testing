package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/nuttchai/go-testing/src/api/app"
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	fmt.Println("about to start the application")
	go app.StartApp() // make app.StartApp() runs in the another thread (otherwise it will block the test by listening the port)
	fmt.Println("application started, about to start test cases")
	os.Exit(m.Run())
}
