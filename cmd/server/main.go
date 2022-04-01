package main

import "covid-19-project/cmd/router"

func main() {
	r, err := router.SetupRouter()
	if err != nil {
		panic(err)
	}

	r.Run()
}
