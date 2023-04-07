package main

import (
	"flight_inquiry/router"
)

func main() {
	e := router.Router()
	e.Run(":8080")
}
