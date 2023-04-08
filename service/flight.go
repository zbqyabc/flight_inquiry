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

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": utils.GetNewSortArr(models.FlightBody.FlightList),
	})
}
