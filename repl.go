package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokeapi"
	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type Pokedex map[string]pokeapi.Pokemon

func startRepl(c *config, cache pokecache.Cache) {
	userInput := bufio.NewScanner(os.Stdin)
	Pokedex := make(Pokedex)

	for {
		fmt.Print("Pokedex > ")
		scan := userInput.Scan()
		if !scan {
			fmt.Println("incorrect input")
			break
		}
		inputStr := userInput.Text()

		allCommands := getCommands()

		commandString, commandArgs := cleanInput(inputStr)

		command, ok := allCommands[commandString]

		if ok {
			command.callback(c, cache, Pokedex, commandArgs)
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, pokecache.Cache, Pokedex, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCommand,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays 20 location areas in the Pokemon world, use it again to get the next page",
			callback:    MapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes to the previous page",
			callback:    MapBackCommand,
		},
		"explore": {
			name:        "explore",
			description: "Explores pokemons in a given area",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch",
			description: "Once you explore an area, try to catch a pokemon...",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon you caught",
			callback:    inspectCommand,
		},
	}
}

func cleanInput(text string) (string, string) {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)

	return words[0], strings.Join(words[1:], "-")
}
