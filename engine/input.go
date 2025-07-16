package engine

import (
	"io"
	"log"
	"os"
	"strings"
)

const (
	SpacebarPress   = "input_spacebar_press"
	SpacebarRelease = "input_spacebar_release"
)

func ReadInput() []string {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(data), "")
}
