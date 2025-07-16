package main

import (
	"encoding/json"
	"fmt"
	"os"

	"go3d/engine"
)

// Builds a go-3D project into a binary form and stores it
// into the output directory defined by settings.json
//
// Uses the values of settings.json to determine how the app
// compiles. This also combines the user written code with the
// pre-written engine code to create a fully distrobutable app
// for all go supported operating systems.
//
// Still need to figure out how to interface with main.go in project
// and turn it into a usable executable
func Build() {
	settings := BuildSettings{}

	data, err := os.ReadFile("./settings.json")
	if err != nil {
		fmt.Println("You are not in a go-3D project. Please navigate to a go-3D project or create one using the new command to proceed.")
		panic(err)
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		panic(err)
	}
	eng := engine.Engine{}
}
