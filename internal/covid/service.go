package covid

import (
	"covid-19-project/internal/covid/covidclient"
)

func GetCovidSummary() *CovidResponse {
	client := covidclient.NewCovidClient(covidclient.COVID_BASEURL)
	resp, err := client.GetCovidCases()
	if err != nil {
		panic(err)
	}

	mProvinces := make(map[string]int)
	ageGroup := &AgeGroup{}

	for _, v := range resp.Data {
		// Count by provinces
		if _, found := mProvinces[v.Province]; found {
			mProvinces[v.Province]++
		} else {
			// New province
			mProvinces[v.Province] = 1
		}

		// Count by age เroup
		// There are 3 age groups: 0-30, 31-60, and 60+ if the case has no age data, please count as "N/A" group
		if !v.Age.Valid {
			ageGroup.NA++
		} else {
			if v.Age.ValueOrZero() <= 30 {
				ageGroup.ZeroTo30++
			} else if v.Age.ValueOrZero() <= 60 {
				ageGroup.ThirtyOneTo60++
			} else if v.Age.ValueOrZero() > 60 {
				ageGroup.SixtyPlus++
			}
		}
	}

	return &CovidResponse{Province: mProvinces, AgeGroup: *ageGroup}
}
