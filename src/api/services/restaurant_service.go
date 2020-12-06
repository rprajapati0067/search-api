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
	sortedThree [3]domain.Restaurant
)

func GetRestaurantService(at string, cat string) (*domain.Response, *errors.RestError) {
	var response domain.Response
	apiResponse, apiErr := providers.GetRestaurant(at, cat)
	if apiErr != nil {
		log.Println(fmt.Sprintf("error when trying to access api: %s", apiErr))
		return nil, &errors.RestError{
			Message: "Failed to retrieve data",
			Status:  400,
			Error:   apiErr.Error,
		}
	}
	nearestThreeRest := getNearestThreeRestaurant(apiResponse.Result.Items)
	response.Restaurant = nearestThreeRest
	return &response, nil
}

func getNearestThreeRestaurant(items Restaurants) *[3]domain.Restaurant {
	if items.Len() > 0 {
		sort.Sort(items)

		for i := 0; i < 3; i++ {
			sortedThree[i] = items[i]
		}
	}
	return &sortedThree
}

type Restaurants []domain.Restaurant

func (s Restaurants) Len() int {
	return len(s)
}

func (s Restaurants) Less(i, j int) bool {
	return s[i].Distance < s[j].Distance
}

func (s Restaurants) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
