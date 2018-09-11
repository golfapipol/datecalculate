package calculate

import "time"

func NewDate(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func DurationBetweenDate(startDate, endDate time.Time) int {
	diff := endDate.Sub(startDate)

	durationDate := (diff.Hours() / 24)
	return int(durationDate)
}
