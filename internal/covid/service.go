//go:generate mockgen -source=service.go -destination=mock/mock_service.go -package=mock

package covid

import (
	"sync"

	"gopkg.in/guregu/null.v4"
)

var (
	covidService     *CovidService
	covidServiceOnce sync.Once
)

type CovidService interface {
	GetCovidSummary() (*CovidResponse, error)
}

type CovidServiceImp struct {
	Client CovidClient
}

func ProviderCovidService(c CovidClient) *CovidServiceImp {
	return &CovidServiceImp{
		Client: c,
	}
}

func (s CovidServiceImp) GetCovidSummary() (*CovidResponse, error) {
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
