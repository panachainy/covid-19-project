package main

import "covid-19-project/cmd/router"

func main() {
	r := router.SetupRouter()
	r.Run()
}
