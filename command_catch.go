package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("mention a pokemon to catch")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemonInfo(name)
	if err != nil {
		return errors.New("pokemon not found")
	}

	res := rand.Intn(pokemon.Base_experience)

	fmt.Printf("Throwing a pokeball at %s...\n", pokemon.Name)
	if res > 50 {
		fmt.Printf("You failed to catch %s\n", pokemon.Name)
		return nil
	}
	fmt.Printf("You caught %s\n", pokemon.Name)
	cfg.mapPokemon[pokemon.Name] = pokemon
	return nil
}
