package utils

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

var Ans [][]string
var path []string

var Citylist = []string{"A", "B", "c", "D"}

func GetCityIndexByName(name string) int {
	for idx, val := range Citylist {
		if val == name {
			return idx
		}
	}
	return -1
}

func removeIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// Dijkstra邻接矩阵算法
func Dfs(graph [][]int, start int, end int) {
	g_length := len(graph)
	visit := make([]int, g_length)
	visit[start] = 1
	path = append(path, getCityName(start))
	if start == end {
		cPath := make([]string, len(path))
		copy(cPath, path)
		Ans = append(Ans, cPath)
	} else {
		for i := 0; i < g_length; i++ {
			if visit[i] == 0 && i != start && graph[start][i] == 1 {
				Dfs(graph, i, end)
			}
		}
	}

	path = removeIndex(path, len(path)-1)
	visit[start] = 0

}
