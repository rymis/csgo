package main

import (
	"os"
	"fmt"
	"github.com/rymis/csgo/csgo"
)

func main() {
	cube := csgo.NewCube(10.0, 10.0, 10.0)
	sphere := csgo.NewSphere(6)

	res := cube.Subtract(sphere)

	f, err := os.Create("test.stl")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	csgo.SaveStl(f, res)
}
