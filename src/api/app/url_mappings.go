package app

import "github.com/rprajapati0067/search-api/src/api/controllers"

func mapsUrls() {
	router.GET("/ping", controllers.Ping)
	router.GET("/search", controllers.GetRestaurant)

}
