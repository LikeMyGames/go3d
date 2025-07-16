package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type (

	// The type that describes the format of a command that can be used
	// with the go3d command
	//
	// Example:
	//  go3d <command-name> <command-args ...>
	//
	// When the command is run, the vale of [Action] is run as a func
	// When the help command is used, the short description is displayed
	// When the help <command-name ...> command is used, the long description is displayed
	Command struct {
		Action           func()
		ShortDescription string
		LongDescription  string
	}

	// The type that describes the settings.json file
	// at the root of a go-3D project
	BuildSettings struct {
		Name           string          `json:"name"`            // The name of a go-3D project (this gets used to generate file names when building)
		BuildOutput    string          `json:"build_output"`    // The build output when building a go-3D project (this defaults to [./bin])
		Entrance       string          `json:"entrance"`        // The entry file into the go program (base file of the app, [./src/main.go] by default)
		CompiledMeshes bool            `json:"compiled_meshes"` // Whether or not meshes get compiled into go-structs when the app gets built (off by default)
		Scripts        SettingsScripts `json:"scripts"`         // Scripts that can be attached to a go-3D project to add extra command line functionality
		Icon           string          `json:"icon"`            // The file location of a .ico file that will become the icon of the app
		Splash         string          `json:"spash"`           // The file location of a png file that will become the splash opening image of the app
		Opening        string          `json:"opening"`         // the file location of a opening .mp4 video that plays once the app has loaded
	}

	SettingsScripts map[string]string // A key-value pair type that defines scripts in the settings.json file of a go-3D project
)

func main() {
	commands := map[string]Command{
		"build": {
			Action:           func() { Build() },
			ShortDescription: "A command that builds a go-3D project into an executable, distrobutable app.",
			LongDescription:  "Uses the settings.json file of a go-3D project to build the project into a distrobutable app.",
		},
		"new": {
			Action:           func() { GenerateNewProject(os.Args[2]) },
			ShortDescription: "A command that sets up a new go-3D project that can have other commands run on it.",
			LongDescription:  "A command that sets up a new go-3D project with a settings.json file (the settings for the project) and a main.go file (the main file for the project).",
		},
		"run": {
			Action:           func() { Run(os.Args[2]) },
			ShortDescription: "A command that can run scripts defined in a go-3D project",
			LongDescription:  "A command that can run scripts defined in a go-3D project's settings.json file. This scripts are defined in the script object's key-value pairs.",
		},
	}

	if os.Args[1] == "help" {
		if len(os.Args[1:]) == 1 {
			for i, v := range commands {
				fmt.Printf("%s: %s\n", i, v.ShortDescription)
			}
		} else {
			for _, v := range os.Args[2:] {
				fmt.Printf("%s:\n%s\n", v, commands[v].LongDescription)
			}
		}
	} else {
		commands[os.Args[1]].Action()
	}
}

// Creates the base for a go-3D project, which includes:
//
// 1. A settings.json file that holds all build settings used in [Build()]
//
// 2. A file and folder structure that is used by the builder for input and output
//
// 3. A base [main.go] file in the [./src] directory that contains boilerplate code
//
// Example:
//
//	GenerateNewProject("test-project")
//
// results in a heirarchy that resembles the following
//
//	 > test-project
//		> bin
//		> public
//			icon.png
//			splash.png
//			opening.mp4
//		> src
//			main.go
//		settings.json
func GenerateNewProject(name string) {
	DefaultSettings := BuildSettings{
		Name:           name,
		BuildOutput:    "./bin",
		Entrance:       "main.go",
		CompiledMeshes: false,
		Scripts: map[string]string{
			"build": "go3d build main.go",
		},
		Icon:    "./public/icon.ico",
		Splash:  "./public/splash.png",
		Opening: "./public/opening.mp4",
	}

	Boilerplate := map[string]string{
		"main.go": `package game

import (
	eng "github.com/likemygames/go3D/engine"
)

func Game() eng.Game {
	return eng.Game{
		Name: "Game Name",
		Levels: []eng.Level{
			{
				Name:   "Level 1",
				Player: eng.Player{},
				Objects: []eng.Object{
					{
						Name:    "Default Cube",
						Variant: eng.MeshType,
						Mesh: eng.Mesh{
							File: "cube.obj",
							Name: "Cube",
						},
						Children: []eng.Object{},
					},
				},
			},
		},
	}
}
`,
	}

	err := os.Mkdir(name, os.ModeDir)
	if err != nil {
		panic(err)
	}

	err = os.Chdir(fmt.Sprintf("./%s", name))
	if err != nil {
		panic(err)
	}

	file, err := os.Create("settings.json")
	if err != nil {
		panic(err)
	}
	data, err := json.MarshalIndent(DefaultSettings, "", "\t")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
	file.Close()

	err = os.Mkdir("bin", os.ModeDir)
	if err != nil {
		panic(err)
	}

	err = os.Mkdir("public", os.ModeDir)
	if err != nil {
		panic(err)
	}

	os.Create("./public/icon.ico")
	os.Create("./public/splash.png")
	os.Create("./public/opening.mp4")

	err = os.Mkdir("src", os.ModeDir)
	if err != nil {
		panic(err)
	}

	file, err = os.Create("./src/main.go")
	if err != nil {
		panic(err)
	}
	_, err = file.Write([]byte(Boilerplate["main.go"]))
	if err != nil {
		panic(err)
	}
}

// Runs scripts defined in a go-3D project's settings.json file
//
// Example:
//
//	Run("build")
//
//	 "scripts": {
//	 	"build": "go3d build" // runs this script
//	 	"test": ...
//	 }
func Run(script string) {
	settings := BuildSettings{}

	data, err := os.ReadFile("./settings.json")
	if err != nil {
		fmt.Println("You are not in a go-3D project. Please navigate to a go-3D project or create one using \"go3d new <project name>\" in order to proceed")
		panic(err)
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		panic(err)
	}
	cmd := strings.Split(settings.Scripts[script], " ")
	c := exec.Command(cmd[0], cmd[1:]...)
	c.Stdin = os.Stdin
	c.Stderr = os.Stderr
	err = c.Run()
	if err != nil {
		panic(err)
	}
}
