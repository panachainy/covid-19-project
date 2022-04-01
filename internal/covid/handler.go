package covid

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	covidHandler     *CovidHandlerImp
	covidHandlerOnce sync.Once
)

type CovidHandler interface {
	GetCovidSummary(c *gin.Context)
}

type CovidHandlerImp struct {
	Service CovidService
}

func ProviderCovidHandler(s CovidService) *CovidHandlerImp {
	covidHandlerOnce.Do(func() {
		covidHandler = &CovidHandlerImp{
			Service: s,
		}
	})

	return covidHandler
}

// /covid/summary
func (h CovidHandlerImp) GetCovidSummary(c *gin.Context) {
	result, err := h.Service.GetCovidSummary()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, result)
}
