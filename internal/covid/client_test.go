package covid_test

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"covid-19-project/internal/covid"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"gopkg.in/guregu/null.v4"
)

var client = resty.New()

func setupTestSuite(tb testing.TB) func(tb testing.TB) {
	// Get the underlying HTTP Client and set it to Mock
	httpmock.ActivateNonDefault(client.GetClient())

	return func(tb testing.TB) {
		fmt.Println("Teardown suite test")
		httpmock.DeactivateAndReset()
	}
}

func TestCovidClientImp_GetCovidCases(t *testing.T) {
	teardownSuite := setupTestSuite(t)
	defer teardownSuite(t)

	type fields struct {
		Client func() *resty.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    *covid.Covid19
		wantErr bool
	}{
		{
			name: "when_call_normally_should_success",
			fields: fields{
				Client: func() *resty.Client {
					// remove old mock
					httpmock.Reset()

					fakeUrl := "/devinterview/covid-cases.json"
					fixture := &covid.Covid19{Data: []covid.Covid19Data{{Province: null.StringFrom("xxx"), Age: null.IntFrom(2)}}}

					mockResponder, err := httpmock.NewJsonResponder(200, fixture)
					if err != nil {
						t.Fatalf("fixture is invalid")
					}

					httpmock.RegisterResponder("GET", fakeUrl, mockResponder)

					return client
				},
			},
			want:    &covid.Covid19{Data: []covid.Covid19Data{{ConfirmDate: "", No: interface{}(nil), Age: null.Int{NullInt64: sql.NullInt64{Int64: 2, Valid: true}}, Gender: "", GenderEn: "", Nation: interface{}(nil), NationEn: "", Province: null.String{NullString: sql.NullString{String: "xxx", Valid: true}}, ProvinceID: 0, District: interface{}(nil), ProvinceEn: "", StatQuarantine: 0}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := covid.CovidClientImp{
				Client: tt.fields.Client(),
			}
			got, err := c.GetCovidCases()
			if (err != nil) != tt.wantErr {
				t.Errorf("CovidClientImp.GetCovidCases() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CovidClientImp.GetCovidCases() = %v, want %v", got, tt.want)
			}
		})
	}
}
