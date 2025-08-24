package main

import "github.com/rymis/csgo/csgo"

// 3D object
type ThreeConverter interface {
	ToModel() *csgo.Model
}

type Cube struct {


}
