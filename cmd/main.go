package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Year struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func main() {
	http.HandleFunc("/time/RFC3339", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().UTC()

		// JSON object creation
		timeData := Year{
			DayOfWeek:  currentTime.Weekday().String(),
			DayOfMonth: currentTime.Day(),
			Month:      currentTime.Month().String(),
			Year:       currentTime.Year(),
			Hour:       currentTime.Hour(),
			Minute:     currentTime.Minute(),
			Second:     currentTime.Second(),
		}

		jsonData, err := json.Marshal(timeData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
