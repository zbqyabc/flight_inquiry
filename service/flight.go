package service

import (
	"flight_inquiry/models"
	"flight_inquiry/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FlightList(c *gin.Context) {
	if err := c.BindJSON(&models.FlightBody); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Param Error!",
		})
		log.Fatal(err)
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
