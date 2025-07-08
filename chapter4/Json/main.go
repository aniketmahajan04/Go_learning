package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"` // omitempty does not include false values
	Actors []string // Only exported fields are marshaled to JSON
}

func main() {
	movies := []Movie{
		{
			Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title: "Kaithi", Year: 2019, Color: true,
			Actors: []string{"Karthi", "Narain"},
		},
		{
			Title: "Vikram", Year: 2022, Color: true,
			Actors: []string{"Kamal Hasan", "Vijay Sethupathi"},
		},
	}
	// Conveting a Go data structure to JSON format is called marshaling.
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatal("JSON marshaling failed:", err)
	}
	fmt.Printf("%s\n", data)
}
