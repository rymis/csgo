package csgo

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Alias for mgl32.Vec3
type Vec3 = mgl32.Vec3

// Alias for mgl32.Vec2
type Vec2 = mgl32.Vec2

// Alias for mgl32.Mat3
type Mat3 = mgl32.Mat3

// Alias for mgl32.Mat4
type Mat4 = mgl32.Mat4

// Constant used in several places to compare floating point numbers.
const Epsilon float32 = 1e-5
