//go:build wireinject
// +build wireinject

//go:generate wire
package covid

import (
	"covid-19-project/internal/covid/covidclient"
	"covid-19-project/internal/covid/covidhandler"
	"covid-19-project/internal/covid/covidservice"

	"github.com/google/wire"
)

func Wire() (covidhandler.CovidHandler, error) {
	wire.Build(covidservice.ProviderCovidService, covidclient.ProviderCovidClient, covidhandler.ProviderCovidHandler,
		wire.Bind(new(covidclient.CovidClient), new(*covidclient.CovidClientImp)),
		wire.Bind(new(covidservice.CovidService), new(*covidservice.CovidServiceImp)),
		wire.Bind(new(covidhandler.CovidHandler), new(*covidhandler.CovidHandlerImp)),
	)
	return covidhandler.CovidHandlerImp{}, nil
}
