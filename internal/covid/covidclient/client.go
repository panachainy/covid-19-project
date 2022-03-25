package covidclient

import (
	"crypto/tls"
	"sync"

	"github.com/go-resty/resty/v2"
)

// TODO: make it support ENV
var (
	COVID_BASEURL   = "http://static.wongnai.com"
	COVID_CASE_PATH = "/devinterview/covid-cases.json"
)

var (
	lock        = &sync.Mutex{}
	covidClient CovidClient
)

type CovidClient interface {
	GetCovidCases() (*Covid19, error)
}

type covidClientImp struct {
	Client *resty.Client
}

func NewCovidClient(baseUrl string) CovidClient {
	if covidClient != nil {
		return covidClient
	}

	lock.Lock()
	defer lock.Unlock()

	client := resty.New()

	// FIXME: Currently WN API can't get with secure way because cert is invalid
	// after fix that API remove this line
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	client.SetBaseURL(baseUrl)

	covidClient = &covidClientImp{Client: client}
	return covidClient
}

func (c covidClientImp) GetCovidCases() (*Covid19, error) {
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
