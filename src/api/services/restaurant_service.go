package services

import (
	"fmt"
	"sort"

	"github.com/rprajapati0067/search-api/src/api/domain"
	"github.com/rprajapati0067/search-api/src/api/providers"
	"github.com/rprajapati0067/search-api/src/api/utils/errors"
	"log"
)

var (
	sortedThree [3]domain.Common
	response    *domain.Response
	apiResponse *domain.Data
	apiErr      *errors.RestError
)

const (
	searchRestaurant    = "restaurant"
	searchTransport     = "transport"
	searchPetrolStation = "petrol-station"
)

type Commons []domain.Common

func (s Commons) Len() int {
	return len(s)
}

func (s Commons) Less(i, j int) bool {
	return s[i].Distance < s[j].Distance
}

func (s Commons) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func getSearchData(at string, cat string, searchData string) (*domain.Response, *errors.RestError) {

	if searchData == searchRestaurant {
		apiResponse, apiErr = providers.GetRestaurant(at, cat)
	} else if searchData == searchTransport {
		apiResponse, apiErr = providers.GetTransport(at, cat)
	} else if searchData == searchPetrolStation {
		apiResponse, apiErr = providers.GetPetrolStation(at, cat)
	}

	if apiErr != nil {
		log.Println(fmt.Sprintf("error when trying to access api: %s", apiErr))
		return nil, &errors.RestError{
			Message: "Failed to retrieve data",
			Status:  400,
			Error:   apiErr.Error,
		}
	}
	return response, nil

}

func getNearestThreeRestaurant(items Commons) *[3]domain.Common {
	if items.Len() > 0 {
		sort.Sort(items)

		for i := 0; i < 3; i++ {
			sortedThree[i] = items[i]
		}
	}
	return &sortedThree
}

func GetRestaurantService(at string, cat string) (*domain.Response, *errors.RestError) {
	response = new(domain.Response)

	result, _ := getSearchData(at, cat, searchRestaurant)
	nearestThreeRest := getNearestThreeRestaurant(apiResponse.Result.Items)
	response.Restaurant = nearestThreeRest
	return result, nil
}

func GetTransportService(at string, cat string) (*domain.Response, *errors.RestError) {
	response = new(domain.Response)

	result, _ := getSearchData(at, cat, searchTransport)
	nearestThreeRest := getNearestThreeRestaurant(apiResponse.Result.Items)
	response.Transport = nearestThreeRest
	return result, nil
}

func GetPetrolPumpService(at string, cat string) (*domain.Response, *errors.RestError) {
	response = new(domain.Response)

	result, _ := getSearchData(at, cat, searchPetrolStation)
	nearestThreeRest := getNearestThreeRestaurant(apiResponse.Result.Items)
	response.PetrolStation = nearestThreeRest
	return result, nil
}
