package covidhandler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"covid-19-project/internal/covid/covidservice"
	"covid-19-project/internal/covid/covidservice/mockcovidservice"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetSummaryHandler(t *testing.T) {
	type args struct {
		service func(ctrl *gomock.Controller) covidservice.CovidService
	}
	type want struct {
		statusCode int
		body       string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "when_call_normal_should_success",
			args: args{
				service: func(ctrl *gomock.Controller) covidservice.CovidService {
					mock := mockcovidservice.NewMockCovidService(ctrl)

					mock.EXPECT().GetCovidSummary().Return(
						&covidservice.CovidResponse{
							Province: map[string]int{"Amnat Charoen": 17},
							AgeGroup: covidservice.AgeGroup{ZeroTo30: 4, ThirtyOneTo60: 10, SixtyPlus: 2, NA: 1},
						},
						nil)

					return mock
				},
			},
			want: want{
				statusCode: 200,
				body:       `{"AgeGroup": {"0-30":4, "31-60":10, "61+":2, "N/A":1}, "Province": {"Amnat Charoen":17}}`,
			},
		},
		{
			name: "when_service_error_should_error",
			args: args{
				service: func(ctrl *gomock.Controller) covidservice.CovidService {
					mock := mockcovidservice.NewMockCovidService(ctrl)

					mock.EXPECT().GetCovidSummary().Return(nil, fmt.Errorf("error from service ja"))

					return mock
				},
			},
			want: want{
				statusCode: 500,
				body:       `{"message":"error from service ja"}`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			summaryHandler := GetSummaryHandler(tt.args.service(ctrl))

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = &http.Request{
				Header: make(http.Header),
			}

			summaryHandler(c)

			got := string(w.Body.Bytes())

			assert.Equal(t, tt.want.statusCode, w.Code, tt.name)
			assert.JSONEq(t, tt.want.body, got, tt.name)
		})
	}
}
