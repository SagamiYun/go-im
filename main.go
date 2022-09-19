package main

import "im/router"

func main() {
	e := router.Router()
	err := e.Run(":8080")
	if err != nil {
		return
	}
}
