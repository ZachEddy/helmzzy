package main

import (
	"fmt"
	"sort"

	"github.com/ZachEddy/helmzzy/command"
)

// Usage returns a message to help people use helmzzy
func Usage() {
	fmt.Println(usage())
}

func usage() string {
	message :=
		`Helmzzy is a tool for managing Helm releases with fzf

Usage: "helmzzy [command]"

Commands:

`
	commandNames := []string{}
	// putting keys into an array and sorting will guarantee the ordering
	for commandName := range command.CommandMap {
		commandNames = append(commandNames, commandName)
	}
	sort.Strings(commandNames)
	for _, commandName := range commandNames {
		message += fmt.Sprintf("    %s\t%s\n",
			commandName, command.CommandMap[commandName].Describe(),
		)
	}
	return message
}
