package main

import (
	"covid-19-project/cmd/router"
	"covid-19-project/internal/covid"
)

func main() {
	covidHandler, err := covid.Wire()
	if err != nil {
		panic(err)
	}

	r := router.SetupRouter(covidHandler)

	r.Run()
}
