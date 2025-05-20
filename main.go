package main

import (
	"thtml/commands"
	"thtml/config"
)

func main() {
	err := config.Start()
	if err != nil {
		panic(err.Error())
	}
	err = commands.HandleCommands()
	if err != nil {
		panic(err.Error())
	}
}
