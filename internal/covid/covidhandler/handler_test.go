package covidhandler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"covid-19-project/internal/covid/covidservice"
	"covid-19-project/internal/covid/covidservice/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCovidHandlerImp_GetCovidSummary(t *testing.T) {
	type want struct {
		statusCode int
		body       string
	}
	type fields struct {
		Service func(ctrl *gomock.Controller) covidservice.CovidService
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "when_normal_should_success",
			fields: fields{
				Service: func(ctrl *gomock.Controller) covidservice.CovidService {
					mock := mock.NewMockCovidService(ctrl)

					mock.EXPECT().GetCovidSummary().Return(
						&covidservice.CovidResponse{
							Province: map[string]int{"Amnat Charoen": 17},
							AgeGroup: map[string]int{"0-30": 4, "31-60": 2, "61+": 1, "N/A": 3},
						},
						nil)

					return mock
				},
			},
			want: want{
				statusCode: 200,
				body:       `{"AgeGroup": {"0-30":4, "31-60":2, "61+":1, "N/A":3}, "Province": {"Amnat Charoen":17}}`,
			},
		},
		{
			name: "when_service_error_should_error",
			fields: fields{
				Service: func(ctrl *gomock.Controller) covidservice.CovidService {
					mock := mock.NewMockCovidService(ctrl)
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
		gin.SetMode(gin.TestMode)

		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			h := CovidHandlerImp{
				Service: tt.fields.Service(ctrl),
			}
			w := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(w)

			c.Request = &http.Request{
				Header: make(http.Header),
			}

			h.GetCovidSummary(c)

			got := string(w.Body.Bytes())

			assert.Equal(t, tt.want.statusCode, w.Code, tt.name)
			assert.JSONEq(t, tt.want.body, got, tt.name)
		})
	}
}
