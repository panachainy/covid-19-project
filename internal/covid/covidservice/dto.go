package covidservice

type AgeGroup struct {
	ZeroTo30      int `json:"0-30"`
	ThirtyOneTo60 int `json:"31-60"`
	SixtyPlus     int `json:"61+"`
	NA            int `json:"N/A"`
}

// TODO: move to stick with handler
type CovidResponse struct {
	Province map[string]int `json:"Province"`
	AgeGroup AgeGroup       `json:"AgeGroup"`
}
