package covid

import "gopkg.in/guregu/null.v4"

type Covid19 struct {
	Data []Covid19Data `json:"Data"`
}

type Covid19Data struct {
	ConfirmDate    string      `json:"ConfirmDate"`
	No             interface{} `json:"No"`
	Age            null.Int    `json:"Age"`
	Gender         string      `json:"Gender"`
	GenderEn       string      `json:"GenderEn"`
	Nation         interface{} `json:"Nation"`
	NationEn       string      `json:"NationEn"`
	Province       null.String `json:"Province"`
	ProvinceID     int         `json:"ProvinceId"`
	District       interface{} `json:"District"`
	ProvinceEn     string      `json:"ProvinceEn"`
	StatQuarantine int         `json:"StatQuarantine"`
}
type CovidResponse struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}
