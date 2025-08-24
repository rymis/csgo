package csgo

import (
	"math"
)

// Create centered cube with sides sx, sy, sz.
func NewCube(sx, sy, sz float32) *Model {
	mesh := NewModel()

	points := []Vec3{
		{-0.5, 0.5, 0.5},
		{0.5, -0.5, 0.5},
		{0.5, 0.5, 0.5},
		{-0.5, -0.5, 0.5},
		{-0.5, -0.5, -0.5},
		{0.5, 0.5, -0.5},
		{0.5, -0.5, -0.5},
		{-0.5, 0.5, -0.5},
	}

	scale := func(v Vec3) Vec3 {
		return Vec3{v.X() * sx, v.Y() * sy, v.Z() * sz}
	}

	polygon := func(i1, i2, i3 int) {
		mesh.AddFace(NewFace(scale(points[i1-1]), scale(points[i2-1]), scale(points[i3-1])))
	}

	polygon(1, 2, 3)
	polygon(2, 1, 4)
	polygon(5, 6, 7)
	polygon(6, 5, 8)
	polygon(5, 2, 4)
	polygon(2, 5, 7)
	polygon(2, 6, 3)
	polygon(6, 2, 7)
	polygon(6, 1, 3)
	polygon(1, 6, 8)
	polygon(5, 1, 8)
	polygon(1, 5, 4)

	return mesh
}

// Create new Sphere. If specified the second parameter defines number of slices.
// If specified the third parameter defines number of stacks.
func NewSphere(r float32, params ...int) *Model {
	res := NewModel()
	slices := 16

	if len(params) > 0 {
		slices = params[0]
	}

	stacks := slices / 2
	if len(params) > 1 {
		stacks = params[1]
	}

	vertex := func(dst []Vertex, theta_in, phi_in float32) []Vertex {
		theta := float64(theta_in) * math.Pi * 2.0
		phi := float64(phi_in) * math.Pi
		dir := Vec3{float32(math.Cos(theta) * math.Sin(phi)), float32(math.Cos(phi)), float32(math.Sin(theta) * math.Sin(phi))}
		uv := Vec2{float32(theta), float32(phi)}
		return append(dst, Vertex{dir.Mul(r), dir, uv})
	}

	for i := 0; i < slices; i++ {
		for j := 0; j < stacks; j++ {
			vertices := make([]Vertex, 0)

			vertices = vertex(vertices, float32(i)/float32(slices), float32(j)/float32(stacks))
			if j > 0 {
				vertices = vertex(vertices, float32(i+1)/float32(slices), float32(j)/float32(stacks))
			}
			if j < stacks-1 {
				vertices = vertex(vertices, float32(i+1)/float32(slices), float32(j+1)/float32(stacks))
			}
			vertices = vertex(vertices, float32(i)/float32(slices), float32(j+1)/float32(stacks))

			for k := 2; k < len(vertices); k++ {
				f := NewFaceWithNormals(
					vertices[0].Position, vertices[0].Normal,
					vertices[k-1].Position, vertices[k-1].Normal,
					vertices[k].Position, vertices[k].Normal)
				f.Vertices[0].UV = vertices[0].UV
				f.Vertices[1].UV = vertices[k-1].UV
				f.Vertices[2].UV = vertices[k].UV

				res.AddFace(f)
			}
		}
	}

	return res
}

// Create new cylinder. If slices not specified it defaults to 16
func NewCylinder(r1, r2, h float32, slices_in ...int) *Model {
	mesh := NewModel()

	slices := 16
	if len(slices_in) > 0 {
		slices = slices_in[0]
	}

	if h <= 0.0 {
		h = 1.0
	}

	circle := func(i int, z float32) Vec3 {
		a := (2.0 * math.Pi * float64(i)) / float64(slices)
		r := r1 + (r2-r1)*z/h
		return Vec3{r * float32(math.Cos(a)), r * float32(math.Sin(a)), 0.0}
	}

	// Add top and bottom:
	for i := 0; i < slices; i++ {
		f := NewFaceWithNormals(
			circle((i+1)%slices, 0.0), Vec3{0.0, 0.0, -1.0},
			circle(i, 0.0), Vec3{0.0, 0.0, -1.0},
			Vec3{0.0, 0.0, 0.0}, Vec3{0.0, 0.0, -1.0})
		f.Vertices[0].UV = Vec2{0.5, 0.5}.Add(f.Vertices[0].Position.Vec2().Mul(0.5 / r2))
		f.Vertices[1].UV = Vec2{0.5, 0.5}.Add(f.Vertices[1].Position.Vec2().Mul(0.5 / r2))
		f.Vertices[2].UV = Vec2{0.5, 0.5}

		mesh.AddFace(f)
	}

	dh := Vec3{0.0, 0.0, h}
	for i := 0; i < slices; i++ {
		f := NewFaceWithNormals(
			Vec3{0.0, 0.0, h}, Vec3{0.0, 0.0, 1.0},
			circle(i, h).Add(dh), Vec3{0.0, 0.0, 1.0},
			circle((i+1)%slices, h).Add(dh), Vec3{0.0, 0.0, 1.0})
		f.Vertices[2].UV = Vec2{0.5, 0.5}.Add(f.Vertices[2].Position.Vec2().Mul(0.5 / r1))
		f.Vertices[1].UV = Vec2{0.5, 0.5}.Add(f.Vertices[1].Position.Vec2().Mul(0.5 / r1))
		f.Vertices[0].UV = Vec2{0.5, 0.5}

		mesh.AddFace(f)
	}

	for l := 0; l < slices; l++ {
		h1 := Vec3{0.0, 0.0, (float32(l) * h) / float32(slices)}
		h2 := Vec3{0.0, 0.0, ((float32(l) + 1.0) * h) / float32(slices)}

		th1 := float32(l) / float32(slices)
		th2 := float32(l+1) / float32(slices)

		for i := 0; i < slices; i++ {
			j := (i + 1) % slices

			f1 := NewFaceWithNormals(circle(i, th1).Add(h1), circle(i, th1), circle(j, th2).Add(h2), circle(j, th2), circle(i, th2).Add(h2), circle(i, th2))

			f1.Vertices[0].UV = Vec2{th1, float32(i) / float32(slices)}
			f1.Vertices[1].UV = Vec2{th2, float32(i+1) / float32(slices)}
			f1.Vertices[2].UV = Vec2{th2, float32(i) / float32(slices)}

			f2 := NewFaceWithNormals(circle(i, th1).Add(h1), circle(i, th1), circle(j, th1).Add(h1), circle(j, th1), circle(j, th2).Add(h2), circle(j, th2))

			f2.Vertices[0].UV = Vec2{th1, float32(i) / float32(slices)}
			f2.Vertices[1].UV = Vec2{th1, float32(i+1) / float32(slices)}
			f2.Vertices[2].UV = Vec2{th2, float32(i+1) / float32(slices)}

			mesh.AddFace(f1)
			mesh.AddFace(f2)
		}
	}

	return mesh
}
