package api_test

import (
	. "datecal/api"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_ApiCalculateDate_Input_StartDay_1_StartMonth_1_StartYear_2018_EndDay_11_EndMonth_9_EndYear_2018_Should_Be_DurationResponse(t *testing.T) {
	url := "/duration?startDay=1&startMonth=1&startYear=2018&endDay=11&endMonth=9&endYear=2018"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `253 days`

	ApiCalculateDate(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, string(actual))
	}
}
