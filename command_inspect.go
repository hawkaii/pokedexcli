package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("mention a pokemon to Instpect")
	}
	name := args[0]
	pokemon, exists := cfg.mapPokemon[name]

	if exists {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t%s: %d\n", stat.Stat.Name, stat.Base_stat)
		}
		fmt.Printf("Types:\n")
		for _, t := range pokemon.Types {
			fmt.Printf("\t%s\n", t.Type.Name)
		}
		return nil
	}
	fmt.Println("you have not caught that pokemon")

	return nil
}
