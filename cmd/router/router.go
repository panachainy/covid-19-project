package router

import (
	"covid-19-project/internal/covid"

	"github.com/gin-gonic/gin"
)

func SetupRouter() (*gin.Engine, error) {
	r := gin.Default()

	covidHandler, err := covid.Wire()
	if err != nil {
		return nil, err
	}

	r.GET("/covid/summary", covidHandler.GetCovidSummary)
	return r, nil
}
