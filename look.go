package main

import (
	"flight_inquiry/utils"
	"fmt"
)

func getCityName(u int) string {
	var name string
	if u == 0 {
		name = "A"
	} else if u == 1 {
		name = "B"
	} else if u == 2 {
		name = "C"
	} else if u == 3 {
		name = "D"
	}
	return name

}

func removeIndex(s []string, index int) []string {

	return append(s[:index], s[index+1:]...)

}

var ans [][]string
var path []string

// Dijkstra邻接矩阵算法
func dfs(start int, end int) {
	var graph = [][]int{
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 0, 0},
	}
	G_length := len(graph)
	visit := make([]int, G_length)
	visit[start] = 1
	path = append(path, getCityName(start))
	if start == end {
		cPath := make([]string, len(path))
		copy(cPath, path)
		ans = append(ans, cPath)
	} else {
		for i := 0; i < G_length; i++ {
			if visit[i] == 0 && i != start && graph[start][i] == 1 {
				dfs(i, end)
			}
		}
	}

	path = removeIndex(path, len(path)-1)
	visit[start] = 0

}

func main11() {
	start := utils.GetCityIndexByName("A")
	end := utils.GetCityIndexByName("B")

	var graph = [][]int{
		{0, 1, 1, 0},
		{0, 0, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 0, 0},
	}

	utils.Dfs(graph, start, end)
	fmt.Println(utils.Ans)
}
