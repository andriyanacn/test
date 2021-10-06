package main

import (
	"phone_review/config"
	"phone_review/route"
)

//MAIN FUNCTION
func main() {
	config.InitDB()
	e := route.New()
	e.Start(":8080")
}
