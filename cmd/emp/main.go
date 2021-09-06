package main

import "github.com/emp/internal/api"

func main() {
	router := api.RouterSetup()
	router.Run(":8082")

}
