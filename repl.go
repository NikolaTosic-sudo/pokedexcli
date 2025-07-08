package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokeapi"
	"github.com/NikolaTosic-sudo/pokedexcli/internal/pokecache"
	"golang.org/x/term"
)

type config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
}

type Pokedex map[string]pokeapi.Pokemon

func startRepl(c *config, cache pokecache.Cache) {
	// userInput := bufio.NewScanner(os.Stdin)
	// b := make([]byte, 1)

	// Pokedex := make(Pokedex)

	// for {
	// n, _ := os.Stdin.Read()
	// fmt.Print("Dobar ", n, '\n')
	// 	fmt.Print("Pokedex > ")
	// 	scan := userInput.Scan()
	// 	if !scan {
	// 		fmt.Println("incorrect input")
	// 		break
	// 	}
	// 	inputStr := userInput.Text()

	// 	allCommands := getCommands()

	// 	commandString, commandArgs := cleanInput(inputStr)

	// 	command, ok := allCommands[commandString]

	// 	if ok {
	// 		command.callback(c, cache, Pokedex, commandArgs)
	// 	} else {
	// 		fmt.Print("Unknown command\n")
	// 	}
	// }
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Println("Type something (press Ctrl+C to exit):")

	b := make([]byte, 3)

	for {
		n, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println("Error reading from stdin:", err)
			break
		}
		fmt.Print(b, "what is B \n")
		if n > 0 {
			char := rune(b[0])
			char2 := rune(b[1])
			char3 := rune(b[2])
			fmt.Printf("You pressed char 1: %c (char 2: %d, Whatever vhar3 is %v)\n", char, char2, char3)

			if char == 3 {
				fmt.Println("Exiting...")
				break
			}
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
			name:        "explore <location-area>",
			description: "Explores pokemons in a given area",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Once you explore an area, try to catch a pokemon...",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "Inspect a pokemon you caught",
			callback:    inspectCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all pokemons you've caught",
			callback:    pokedexCommand,
		},
	}
}

func cleanInput(text string) (string, string) {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)

	return words[0], strings.Join(words[1:], "-")
}
