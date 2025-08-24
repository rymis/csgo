package main

import (
	"os"
	"fmt"
	"github.com/rymis/csgo/csgo"
)

func saveStl(fnm string, model *csgo.Model) {
	f, err := os.Create(fnm)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	csgo.SaveStl(f, model)
}

func main() {
	cube := csgo.NewCube(10.0, 10.0, 10.0)
	sphere := csgo.NewSphere(6)

	saveStl("subtract.stl", cube.Subtract(sphere))
	saveStl("union.stl", cube.Join(sphere))
	saveStl("intersection.stl", cube.Intersect(sphere))
	saveStl("hull.stl", cube.Hull(sphere))
}
