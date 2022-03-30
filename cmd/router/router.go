package router

import (
	"covid-19-project/internal/covid/covidhandler"
	"covid-19-project/internal/covid/covidservice"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/covid/summary", covidhandler.GetSummaryHandler(covidservice.NewCovidService()))
	return r
}
