package utils

import "fmt"

// 数组中是否包含元素
func isContain(arr []string, item string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}

// 搜索元素中slice中的索引
func searchIndexInArray(arr []string, item string) int {
	for k, v := range arr {
		if v == item {
			return k
		}
	}
	return -1
}

// 递归生成新数组
func recur(list_src [][]string, list_new [][]string, arr_start []string, s string) [][]string {
	k := searchIndexInArray(arr_start, s)
	if k != -1 {
		v := list_src[k]
		list_new = append(list_new, v)
		list_new = recur(list_src, list_new, arr_start, v[1])
	}
	return list_new
}

// 新排序数组
func GetNewSortList(list_src [][]string) [][]string {

	var list_new [][]string //新数组
	var arr_start []string  //出发城市数组
	var arr_end []string    //到达城市数组

	//出发城市和到达城市数组
	for _, v := range list_src {
		arr_start = append(arr_start, v[0])
		arr_end = append(arr_end, v[1])
	}

	for k, v := range arr_start {
		if !isContain(arr_end, v) {
			list_new = append(list_new, list_src[k])
		}
	}
	list_new = recur(list_src, list_new, arr_start, list_new[0][1])
	fmt.Println(list_new)
	return list_new

}
