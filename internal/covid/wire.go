//go:build wireinject
// +build wireinject

//go:generate wire
package covid

import (
	"github.com/google/wire"
)

func Wire() (CovidHandler, error) {
	wire.Build(ProviderCovidService, ProviderCovidClient, ProviderCovidHandler,
		wire.Bind(new(CovidClient), new(*CovidClientImp)),
		wire.Bind(new(CovidService), new(*CovidServiceImp)),
		wire.Bind(new(CovidHandler), new(*CovidHandlerImp)),
	)
	return CovidHandlerImp{}, nil
}
