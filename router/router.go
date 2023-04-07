package router

import (
	"flight_inquiry/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/flight_list", service.FlightList)
	return r
}
