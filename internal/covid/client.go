//go:generate mockgen -source=client.go -destination=mock/mock_client.go -package=mock

package covid

import (
	"crypto/tls"
	"sync"

	"github.com/go-resty/resty/v2"
)

var (
	covidClient     *CovidClientImp
	covidClientOnce sync.Once
)

// TODO: make it support ENV
var (
	COVID_BASEURL   = "http://static.wongnai.com"
	COVID_CASE_PATH = "/devinterview/covid-cases.json"
)

type CovidClient interface {
	GetCovidCases() (*Covid19, error)
}

type CovidClientImp struct {
	Client *resty.Client
}

func ProviderCovidClient() *CovidClientImp {
	client := resty.New()

	// FIXME: Currently WN API can't get with secure way because cert is invalid
	// after fix that API remove this line
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetBaseURL(COVID_BASEURL)

	covidClientOnce.Do(func() {
		covidClient = &CovidClientImp{Client: client}
	})

	return covidClient
}

func (c CovidClientImp) GetCovidCases() (*Covid19, error) {
	resp, err := c.Client.R().
		// EnableTrace().
		SetResult(&Covid19{}).
		Get(COVID_CASE_PATH)
	if err != nil {
		return nil, err
	}

	data := resp.Result().(*Covid19)

	return data, nil
}
