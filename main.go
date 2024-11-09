package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type meteo_status struct {
	Humidity int     `json:"Humidity"`
	Degrees  float32 `json:"Degrees"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var meteo meteo_status
	switch r.URL.Path {
	case "/all":
		if r.Method == "GET" {
			fmt.Println(meteo)
		} else if r.Method == "POST" || r.Method == "PUT" {
			fmt.Println(json.NewDecoder(r.Body).Decode(&meteo))
		}
	case "/hum":
		if r.Method == "GET" {
			fmt.Println(meteo.Humidity)
		} else if r.Method == "POST" || r.Method == "PUT" {
			fmt.Println(json.NewDecoder(r.Body).Decode(&meteo))
		}
	case "/deg":
		if r.Method == "GET" {
			fmt.Println(meteo.Degrees)
		} else if r.Method == "POST" || r.Method == "PUT" {
			fmt.Println(json.NewDecoder(r.Body).Decode(&meteo))
		}
	default:
		fmt.Println("Error!")
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Starting server for testing HTTP requests... ")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
