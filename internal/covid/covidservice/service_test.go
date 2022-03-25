package covidservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"covid-19-project/internal/covid/covidclient"
	"covid-19-project/internal/covid/covidclient/mockcovidclient"

	"github.com/golang/mock/gomock"
)

func Test_covidServiceImp_GetCovidSummary(t *testing.T) {
	type fields struct {
		Client func(ctrl *gomock.Controller) covidclient.CovidClient
	}
	tests := []struct {
		name    string
		fields  fields
		want    *CovidResponse
		wantErr bool
	}{
		{
			name: "when_call_normal_should_success",
			fields: fields{
				Client: func(ctrl *gomock.Controller) covidclient.CovidClient {
					mock := mockcovidclient.NewMockCovidClient(ctrl)

					b, err := ioutil.ReadFile("../covidclient/mockcovidclient/01_covid19_response.json")
					if err != nil {
						t.Fatal(err)
					}

					mockCovid := &covidclient.Covid19{}

					if err := json.Unmarshal(b, mockCovid); err != nil {
						t.Fatal(err)
					}

					mock.EXPECT().GetCovidCases().Return(mockCovid, nil)

					return mock
				},
			},
			want:    &CovidResponse{Province: map[string]int{"Amnat Charoen": 17, "Ang Thong": 36, "Bangkok": 27, "Bueng Kan": 23, "Buriram": 18, "Chachoengsao": 24, "Chai Nat": 25, "Chaiyaphum": 28, "Chanthaburi": 17, "Chiang Mai": 22, "Chiang Rai": 15, "Chonburi": 29, "Chumphon": 25, "Kalasin": 27, "Kamphaeng Phet": 23, "Kanchanaburi": 23, "Khon Kaen": 27, "Krabi": 27, "Lampang": 24, "Lamphun": 25, "Loei": 17, "Lopburi": 19, "Mae Hong Son": 22, "Maha Sarakham": 26, "Mukdahan": 28, "N/A": 27, "Nakhon Nayok": 19, "Nakhon Pathom": 31, "Nakhon Phanom": 24, "Nakhon Ratchasima": 28, "Nakhon Sawan": 24, "Nakhon Si Thammarat": 35, "Nan": 20, "Narathiwat": 22, "Nong Bua Lamphu": 29, "Nong Khai": 27, "Nonthaburi": 29, "Pathum Thani": 30, "Pattani": 27, "Phang Nga": 28, "Phatthalung": 29, "Phayao": 25, "Phetchabun": 33, "Phetchaburi": 26, "Phichit": 21, "Phitsanulok": 24, "Phra Nakhon Si Ayutthaya": 25, "Phrae": 28, "Phuket": 25, "Prachinburi": 19, "Prachuap Khiri Khan": 34, "Ranong": 35, "Ratchaburi": 21, "Rayong": 25, "Roi Et": 25, "Sa Kaeo": 26, "Sakon Nakhon": 42, "Samut Prakan": 31, "Samut Sakhon": 29, "Samut Songkhram": 22, "Saraburi": 26, "Satun": 37, "Sing Buri": 26, "Sisaket": 27, "Songkhla": 24, "Sukhothai": 23, "Suphan Buri": 28, "Surat Thani": 25, "Surin": 24, "Tak": 18, "Trang": 20, "Trat": 25, "Ubon Ratchathani": 23, "Udon Thani": 34, "Uthai Thani": 24, "Uttaradit": 24, "Yala": 27, "Yasothon": 26}, AgeGroup: AgeGroup{ZeroTo30: 602, ThirtyOneTo60: 607, SixtyPlus: 769, NA: 22}},
			wantErr: false,
		},
		{
			name: "when_external_error_should_error",
			fields: fields{
				Client: func(ctrl *gomock.Controller) covidclient.CovidClient {
					mock := mockcovidclient.NewMockCovidClient(ctrl)

					mock.EXPECT().GetCovidCases().Return(&covidclient.Covid19{}, fmt.Errorf("external error ja"))

					return mock
				},
			},
			want:    &CovidResponse{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			s := covidServiceImp{
				Client: tt.fields.Client(ctrl),
			}
			got, err := s.GetCovidSummary()
			if (err != nil) != tt.wantErr {
				t.Errorf("covidServiceImp.GetCovidSummary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("covidServiceImp.GetCovidSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}
