package main

import (
	"datecal/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/duration", api.ApiCalculateDate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
