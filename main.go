package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type config struct {
	Next     string
	Previous string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type LocationArea struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
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
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Goes to the previous page",
			callback:    mapBackCommand,
		},
	}
}

func cleanInput(text string) []string {
	loweredText := strings.ToLower(text)
	words := strings.Fields(loweredText)

	return words
}

func exitCommand(c *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(c *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func mapCommand(c *config) error {

	if c.Next == "" {
		c.Next = "https://pokeapi.co/api/v2/location-area"
	}

	res, err := http.Get(c.Next)

	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("there was an issue with the request")
	}

	defer res.Body.Close()

	var areas LocationArea

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &areas); err != nil {
		return err
	}

	c.Next = areas.Next

	if areas.Previous != nil {
		c.Previous = *areas.Previous
	}

	for _, area := range areas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}

func mapBackCommand(c *config) error {

	if c.Previous == "" {

		fmt.Print("you're on the first page\n")

		return nil
	}

	res, err := http.Get(c.Previous)

	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("there was an issue with the request")
	}

	defer res.Body.Close()

	var areas LocationArea

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &areas); err != nil {
		return err
	}

	c.Next = areas.Next

	if areas.Previous != nil {
		c.Previous = *areas.Previous
	}

	for _, area := range areas.Results {
		fmt.Printf("%s\n", area.Name)
	}

	return nil
}

func main() {

	userInput := bufio.NewScanner(os.Stdin)

	c := &config{}

	for {
		fmt.Print("Pokedex > ")
		scan := userInput.Scan()
		if !scan {
			fmt.Println("incorrect input")
			break
		}
		inputStr := userInput.Text()

		allCommands := getCommands()

		command, ok := allCommands[inputStr]

		if ok {
			command.callback(c)
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
