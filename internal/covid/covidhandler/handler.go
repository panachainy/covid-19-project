package covidhandler

import (
	"net/http"

	"covid-19-project/internal/covid/covidservice"

	"github.com/gin-gonic/gin"
)

// I write this way because if you need to inject some obj to handler you can use with this way.
//  /covid/summary
func GetSummaryHandler(service covidservice.CovidService) func(c *gin.Context) {
	return func(c *gin.Context) {
		result, err := service.GetCovidSummary()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, result)
	}
}
