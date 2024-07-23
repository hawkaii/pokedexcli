package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Location struct {
	Name string `json:"name"`
}

func getRequest(id int) string {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + fmt.Sprintf("%d/", id))
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	location := Location{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		log.Fatal(err)
	}
	return location.Name

}
