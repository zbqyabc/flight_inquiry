package service

import (
	"flight_inquiry/models"
	"flight_inquiry/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func FlightList(c *gin.Context) {
	if err := c.ShouldBindBodyWith(&models.FlightBody, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Param Error",
		})
		return
	}

	newList := utils.GetNewSortArr(models.FlightBody.FlightList)
	var arrFlight []string //航线起始
	arrFlight = append(arrFlight, newList[0][0])
	arrFlight = append(arrFlight, newList[len(newList)-1][1])

	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "Success",
		"newList":   newList,
		"arrFlight": arrFlight,
	})
}
