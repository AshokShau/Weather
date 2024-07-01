package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type WeatherData struct {
	Name string `json:"name"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	_ = godotenv.Load()

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("API_KEY is not set")
		os.Exit(1)
	}

	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		resp, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}(resp.Body)

		var data WeatherData
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Printf("Starting server at port 8080\n" + "http://localhost:8080/weather/london\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
