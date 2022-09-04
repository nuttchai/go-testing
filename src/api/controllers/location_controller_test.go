package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/gin-gonic/gin"
	"github.com/nuttchai/go-testing/src/api/domain/locations"
	"github.com/nuttchai/go-testing/src/api/services"
	"github.com/nuttchai/go-testing/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}

// Integration test with the controller, service, and api
func TestGetCountryNotFound(t *testing.T) {
	// Mock location service method (this won't call an actual service and provider functions but instead of calling a mock function)
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)

	var appErr errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &appErr)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, appErr.Status)
	assert.EqualValues(t, "Country not found", appErr.Message)
}

func TestGetCountryFound(t *testing.T) {
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, response.Code)

	var country locations.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)

	assert.Nil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
