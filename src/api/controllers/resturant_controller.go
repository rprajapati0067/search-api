package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rprajapati0067/search-api/src/api/domain"
	"github.com/rprajapati0067/search-api/src/api/services"
	"github.com/rprajapati0067/search-api/src/api/utils/errors"
	"net/http"
)

var (
	result          *domain.Response
	restaurantError *errors.RestError
)

func GetRestaurant(c *gin.Context) {

	at := c.Query("at")
	cat := c.Query("cat")
	fmt.Println(c.Request.URL)
	if at != "" && cat != "" {
		result, restaurantError = services.GetRestaurantService(at, cat)
		if restaurantError != nil {
			c.JSON(restaurantError.Status, restaurantError)
			return
		}
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "Bad Request"})
	}

}
