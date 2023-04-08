package utils

import "fmt"

// 数组中是否包含元素
func isContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 搜索元素中slice中的索引
func searchArrIndex(items []string, item string) int {
	for k, eachItem := range items {
		if eachItem == item {
			return k
		}
	}
	return -1
}

// 递归生成新数组
func recur(arr [][]string, arr_new [][]string, arr_start []string, s string) [][]string {
	k := searchArrIndex(arr_start, s)
	if k != -1 {
		v := arr[k]
		arr_new = append(arr_new, v)
		arr_new = recur(arr, arr_new, arr_start, v[1])
	}
	return arr_new
}

// 新排序数组
func GetNewSortArr(arr [][]string) [][]string {

	var arr_new [][]string //新数组
	var arr_start []string //出发城市数组
	var arr_end []string   //到达城市数组

	//出发城市和到达城市数组
	for _, v := range arr {
		arr_start = append(arr_start, v[0])
		arr_end = append(arr_end, v[1])
	}

	for k, v := range arr_start {
		if !isContain(arr_end, v) {
			arr_new = append(arr_new, arr[k])
		}
	}
	arr_new = recur(arr, arr_new, arr_start, arr_new[0][1])
	fmt.Println(arr_new)
	return arr_new

}
