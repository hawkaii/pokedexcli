package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	//check if mapPokemon is empty
	if len(cfg.mapPokemon) == 0 {
		return errors.New("you have not caught any pokemon")
	}

	//iterate over mapPokemon and print the names
	for pokemon := range cfg.mapPokemon {
		fmt.Println("- " + pokemon)
	}

	return nil

}
