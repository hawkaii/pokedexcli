package main

import (
	"time"

	"github.com/hawkaii/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		mapPokemon:    map[string]pokeapi.Pokemon{},
		pokeapiClient: *pokeClient,
	}

	startRepl(cfg)
}
