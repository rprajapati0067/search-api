package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rprajapati0067/search-api/src/api/domain"
	"github.com/rprajapati0067/search-api/src/api/services"
	"net/http"
)

var (
	result *domain.Response
)

func GetRestaurant(c *gin.Context) {

	at := c.Query("at")
	cat := c.Query("cat")

	if at != "" && cat != "" {
		result, restaurantError := services.GetRestaurantService(at, cat)
		if restaurantError != nil {
			c.JSON(restaurantError.Status, restaurantError)
			return
		}
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

}

func GetTransport(c *gin.Context) {

	at := c.Query("at")
	cat := c.Query("cat")

	if at != "" && cat != "" {
		result, transportError := services.GetTransportService(at, cat)
		if transportError != nil {
			c.JSON(transportError.Status, transportError)
			return
		}
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

}

func GetPetrolPStation(c *gin.Context) {

	at := c.Query("at")
	cat := c.Query("cat")

	if at != "" && cat != "" {
		result, petrolStationError := services.GetPetrolPumpService(at, cat)
		if petrolStationError != nil {
			c.JSON(petrolStationError.Status, petrolStationError)
			return
		}
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

}
