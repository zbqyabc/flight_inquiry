package models

// 航线参数消息体
var FlightBody struct {
	FlightList [][]string `json:"flight_list"`
}
