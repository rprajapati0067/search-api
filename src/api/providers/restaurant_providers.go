package providers

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"github.com/rprajapati0067/search-api/src/api/domain"
	"github.com/rprajapati0067/search-api/src/api/utils/errors"
	"log"
)

const (
	url    = "https://places.ls.hereapi.com/places/v1/discover/explore"
	apiKey = "GHT-iUolNrr1E3rV10fJHnRI4-hMnQgLCwMYHPD9Yr4"
)

func GetRestaurant(at string, cat string) (*domain.Data, *errors.RestError) {
	var target domain.Data
	client := resty.New()

	resp, err := client.R().SetResult(domain.Restaurant{}).ForceContentType("application/json").SetQueryParams(map[string]string{
		"cat":    cat,
		"at":     at,
		"apiKey": apiKey,
	}).
		EnableTrace().
		Get(url)

	if err != nil {
		log.Println(fmt.Sprintf("error when trying to access api: %s", err.Error()))
		return nil, &errors.RestError{
			Message: "Failed to retrieve data",
			Status:  400,
			Error:   err.Error(),
		}

	}

	unmarshalError := json.Unmarshal(resp.Body(), &target)
	if unmarshalError != nil {
		log.Println(fmt.Sprintf("error when unmarshalling: %s", unmarshalError.Error()))
		return nil, &errors.RestError{
			Message: "unmarshalling failed",
			Status:  500,
			Error:   unmarshalError.Error(),
		}
	}
	return &target, nil

}
