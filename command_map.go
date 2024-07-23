package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationRes.Next
	cfg.prevLocationURL = locationRes.Previous

	for _, loc := range locationRes.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("No previous location")
	}

	locationRes, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationRes.Next
	cfg.prevLocationURL = locationRes.Previous

	for _, loc := range locationRes.Results {
		fmt.Println(loc.Name)
	}
	return nil

}
