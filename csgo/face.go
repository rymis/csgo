package csgo

// Face representation. Face is a triangle with material identifier.
type Face struct {
	MaterialId string
	Vertices   [3]Vertex
}

// Create a Face based on 3 vectors.
func NewFace(a, b, c Vec3) *Face {
	res := &Face{}
	res.MaterialId = ""
	res.Vertices[0].Position = a
	res.Vertices[1].Position = b
	res.Vertices[2].Position = c
	n := res.Normal()
	res.Vertices[0].Normal = n
	res.Vertices[1].Normal = n
	res.Vertices[2].Normal = n
	res.Vertices[0].UV = Vec2{0.0, 0.0}
	res.Vertices[1].UV = Vec2{0.0, 1.0}
	res.Vertices[2].UV = Vec2{1.0, 0.0}

	return res
}

// Create a Face with normals specified for each vertex.
func NewFaceWithNormals(a, na, b, nb, c, nc Vec3) *Face {
	res := &Face{}
	res.MaterialId = ""
	res.Vertices[0].Position = a
	res.Vertices[1].Position = b
	res.Vertices[2].Position = c
	res.Vertices[0].Normal = na
	res.Vertices[1].Normal = nb
	res.Vertices[2].Normal = nc
	res.Vertices[0].UV = Vec2{0.0, 0.0}
	res.Vertices[1].UV = Vec2{0.0, 1.0}
	res.Vertices[2].UV = Vec2{1.0, 0.0}

	return res
}

// Get face normal. This normal is computed based on vertices and does not take into account vertex normals.
func (f *Face) Normal() Vec3 {
	return f.Vertices[1].Position.Sub(f.Vertices[0].Position).Cross(f.Vertices[2].Position.Sub(f.Vertices[0].Position)).Normalize()
}

// Create a full copy of the Face.
func (f *Face) Clone() *Face {
	res := NewFaceWithNormals(f.Vertices[0].Position, f.Vertices[0].Normal, f.Vertices[1].Position, f.Vertices[1].Normal, f.Vertices[2].Position, f.Vertices[2].Normal)
	res.Vertices[0].UV = f.Vertices[0].UV
	res.Vertices[1].UV = f.Vertices[1].UV
	res.Vertices[2].UV = f.Vertices[2].UV

	return res
}
