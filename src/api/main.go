package main

import (
	"fmt"

	"github.com/nuttchai/go-testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("AR")

	fmt.Println(err)
	fmt.Println(country)
}
