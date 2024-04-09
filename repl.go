package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(" > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) == 0 {
			continue
		}

		inputCommand := cleaned[0]
		avaliableCommands := getCommands()
		command, ok := avaliableCommands[inputCommand]

		if !ok {
			fmt.Println("invalid command")
			continue
		}

		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu.",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Lists the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous page of location areas",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore {location_id_or_name}",
			description: "Explores a specific location",
			callback:    callbackExplore,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the program.",
			callback:    callbackExit,
		},
	}
}
