package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FlightList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": "",
	})
}
