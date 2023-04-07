package service

import (
	"flight_inquiry/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FlightList(c *gin.Context) {

	start_city := c.Query("start_city")
	arrive_city := c.Query("arrive_city")

	fmt.Println(start_city, arrive_city)

	if start_city == "" || arrive_city == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "出发城市或者达到城市不能为空",
		})
		return
	}

	start := utils.GetCityIndexByName(start_city)
	end := utils.GetCityIndexByName(arrive_city)

	if start == -1 || end == -1 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "出发城市或者达到城市不存在",
		})
		return
	}

	var graph = [][]int{
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 0, 0},
	}

	utils.Dfs(graph, start, end)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Success",
		"data": utils.Ans,
	})
}
