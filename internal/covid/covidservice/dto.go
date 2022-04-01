package covidservice

// TODO: move to stick with handler
type CovidResponse struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}
