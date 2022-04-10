package router

import (
	"covid-19-project/internal/covid"

	"github.com/gin-gonic/gin"
)

func SetupRouter(covidHandler covid.CovidHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/covid/summary", covidHandler.GetCovidSummary)
	return r
}
