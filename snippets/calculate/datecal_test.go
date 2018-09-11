package calculate_test

import (
	. "datecal/calculate"
	"testing"
	"time"
)

func Test_NewDate_Input_1_1_2018_Should_Be_1_1_2018(t *testing.T) {
	day := 4
	month := 1
	year := 2018
	expected := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	actual := NewDate(day, month, year)

	if expected != actual {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func Test_DurationBetweenDate_Input_1_1_2018_And_11_9_2018_Should_Be_253(t *testing.T) {
	startDate := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2018, 9, 11, 0, 0, 0, 0, time.UTC)
	expected := 253

	actual := DurationBetweenDate(startDate, endDate)

	if expected != actual {
		t.Errorf("should be %d but it is %d ", expected, actual)

	}
}
