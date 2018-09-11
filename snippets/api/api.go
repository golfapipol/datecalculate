package api

import (
	"datecal/calculate"
	"fmt"
	"net/http"
	"strconv"
)

type Duration struct {
	Days string `json:"days"`
}

func ApiCalculateDate(responseWriter http.ResponseWriter, request *http.Request) {
	queryString := request.URL.Query()
	startDay, _ := strconv.Atoi(queryString.Get("startDay"))
	startMonth, _ := strconv.Atoi(queryString.Get("startMonth"))
	startYear, _ := strconv.Atoi(queryString.Get("startYear"))
	endDay, _ := strconv.Atoi(queryString.Get("endDay"))
	endMonth, _ := strconv.Atoi(queryString.Get("endMonth"))
	endYear, _ := strconv.Atoi(queryString.Get("endYear"))
	startDate := calculate.NewDate(startDay, startMonth, startYear)
	endDate := calculate.NewDate(endDay, endMonth, endYear)

	output := fmt.Sprintf("%d days", calculate.DurationBetweenDate(startDate, endDate))
	responseWriter.Write([]byte(output))
}
