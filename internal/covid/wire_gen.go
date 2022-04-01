// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package covid

// Injectors from wire.go:

func Wire() (CovidHandler, error) {
	covidClientImp := ProviderCovidClient()
	covidServiceImp := ProviderCovidService(covidClientImp)
	covidHandlerImp := ProviderCovidHandler(covidServiceImp)
	return covidHandlerImp, nil
}
