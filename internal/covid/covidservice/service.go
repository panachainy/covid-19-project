package covidservice

import (
	"sync"

	"covid-19-project/internal/covid/covidclient"

	"gopkg.in/guregu/null.v4"
)

var (
	lock         = &sync.Mutex{}
	covidService CovidService
)

type CovidService interface {
	GetCovidSummary() (*CovidResponse, error)
}

type covidServiceImp struct {
	Client covidclient.CovidClient
}

func NewCovidService() CovidService {
	if covidService != nil {
		return covidService
	}

	lock.Lock()
	defer lock.Unlock()

	covidService = &covidServiceImp{
		Client: covidclient.NewCovidClient(covidclient.COVID_BASEURL),
	}
	return covidService
}

func (s covidServiceImp) GetCovidSummary() (*CovidResponse, error) {
	resp, err := s.Client.GetCovidCases()
	if err != nil {
		return &CovidResponse{}, err
	}

	mProvinces := make(map[string]int)
	mAgeGroup := make(map[string]int)

	for _, v := range resp.Data {

		// Set null to N/A
		if !v.Province.Valid {
			v.Province = null.StringFrom("N/A")
		}
		mProvinces[v.Province.ValueOrZero()]++

		// Count by age Group
		// There are 3 age groups: 0-30, 31-60, and 60+ if the case has no age data, please count as "N/A" group
		if !v.Age.Valid {
			mAgeGroup["N/A"]++
		} else {
			if v.Age.ValueOrZero() <= 30 {
				mAgeGroup["0-30"]++
			} else if v.Age.ValueOrZero() <= 60 {
				mAgeGroup["31-60"]++
			} else if v.Age.ValueOrZero() > 60 {
				mAgeGroup["61+"]++
			}
		}
	}

	return &CovidResponse{Province: mProvinces, AgeGroup: mAgeGroup}, nil
}
