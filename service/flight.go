package service

import (
	"flight_inquiry/models"
	"flight_inquiry/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// 获取排序后的航线列表
func FlightList(c *gin.Context) {
	if err := c.ShouldBindBodyWith(&models.FlightBody, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Param Error",
		})
		return
	}

	newFlightList := utils.GetNewSortList(models.FlightBody.FlightList)
	var arrFlight []string //航线起始
	arrFlight = append(arrFlight, newFlightList[0][0])
	arrFlight = append(arrFlight, newFlightList[len(newFlightList)-1][1])

	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "Success",
		"newList":   newFlightList,
		"arrFlight": arrFlight,
	})
}
