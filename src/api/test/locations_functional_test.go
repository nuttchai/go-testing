package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/nuttchai/go-testing/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

// functional testing
func TestGetCountriesNotFound(t *testing.T) {
	// mock api, otherwise it will return its default response of rest library (MockUp nil!)
	rest.FlushMockups()
	rest.AddMockups(&rest.Mock{
		URL:          "https://api.mercadolibre.com/countries/AR",
		HTTPMethod:   http.MethodGet,
		RespHTTPCode: http.StatusNotFound,
		RespBody:     `{"message": "Country not found", "error": "not_found", "status": 404, "cause": []}`,
	})

	fmt.Println("about to functional test get countries")
	// NOTE: this will actually call the api to get data (black box testing)
	// IMPORTANT: This will test an api call with the given conditions to check that it can run correctly or not
	// For example, we want to check "http://localhost:8080/locations/countries/AR" path that can call our api or not
	response, err := http.Get("http://localhost:8080/locations/countries/AR")

	assert.Nil(t, err)
	assert.NotNil(t, response)

	bytes, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))

	var apiErr errors.ApiError
	err = json.Unmarshal(bytes, &apiErr)
	assert.Nil(t, err)

	assert.EqualValues(t, http.StatusNotFound, apiErr.Status)
	assert.EqualValues(t, "not_found", apiErr.Error)
	assert.EqualValues(t, "Country not found", apiErr.Message)
}
