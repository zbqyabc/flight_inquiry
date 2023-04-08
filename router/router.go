package router

import (
	"flight_inquiry/service"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	//排序后的新航线
	r.POST("flight/get_new_list", service.FlightList)

	return r
}
