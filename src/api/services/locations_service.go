package services

import (
	"fmt"

	"github.com/nuttchai/go-testing/src/api/domain/locations"
	"github.com/nuttchai/go-testing/src/api/providers/locations_provider"
	"github.com/nuttchai/go-testing/src/api/utils/errors"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationsService locationsServiceInterface
)

// init() is called when the PACKAGE (not the file itself) is initialized
func init() {
	LocationsService = &locationsService{}
}

func (s *locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	fmt.Println("Calling GetCountry Service")
	return locations_provider.GetCountry(countryId)
}
