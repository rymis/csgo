package csgo

// 3D Model representation. Model is a list of faces.
type Model struct {
	Faces []*Face
}

// Create new empty model.
func NewModel() *Model {
	return &Model{}
}

// Add a new face to the model.
func (m *Model) AddFace(face *Face) {
	m.Faces = append(m.Faces, face)
}

// Set material for all faces.
func (m *Model) SetMaterial(mat string) {
	for _, f := range m.Faces {
		f.MaterialId = mat
	}
}

// Create a full clone of the model.
func (m *Model) Clone() *Model {
	res := NewModel()
	for _, f := range m.Faces {
		res.AddFace(f.Clone())
	}

	return res
}

// Move model by vector.
func (m *Model) Move(d Vec3) {
	for _, f := range m.Faces {
		for i := 0; i < 3; i++ {
			f.Vertices[i].Position = f.Vertices[i].Position.Add(d)
		}
	}
}

// Transform the model.
func (m *Model) Transform(a *Mat3) {
	for _, f := range m.Faces {
		for i := 0; i < 3; i++ {
			f.Vertices[i].Position = a.Mul3x1(f.Vertices[i].Position)
			f.Vertices[i].Normal = a.Mul3x1(f.Vertices[i].Normal).Normalize()
		}
	}
}

// Mirror the model over plane.
func (m *Model) Mirror(plane Vec3) {
	n := plane.Normalize()

	t := func(v Vec3) Vec3 {
		return v.Sub(n.Mul(2.0 * v.Dot(n)))
	}

	for _, f := range m.Faces {
		for i := 0; i < 3; i++ {
			f.Vertices[i].Position = t(f.Vertices[i].Position)
			f.Vertices[i].Normal = t(f.Vertices[i].Normal)
		}
	}
}

// Scale the model.
func (m *Model) Scale(s Vec3) {
	mat := Mat3{s.X(), 0.0, 0.0, 0.0, s.Y(), 0.0, 0.0, 0.0, s.Z()}
	m.Transform(&mat)
}

// Apply CSG transformation: join 2 models
func (m *Model) Join(other *Model) *Model {
	csg1 := m.ToCsg()
	csg2 := other.ToCsg()

	csg := csg1.Union(csg2)

	return NewModelFromCsg(csg)
}

// Apply CSG transformation: subtract 2 models
func (m *Model) Subtract(other *Model) *Model {
	csg1 := m.ToCsg()
	csg2 := other.ToCsg()

	csg := csg1.Subtract(csg2)

	return NewModelFromCsg(csg)
}

// Apply CSG transformation: intersect 2 models
func (m *Model) Intersect(other *Model) *Model {
	csg1 := m.ToCsg()
	csg2 := other.ToCsg()

	csg := csg1.Intersect(csg2)

	return NewModelFromCsg(csg)
}

// Generate a convex hull from several models.
func (m *Model) Hull(other *Model) *Model {
	idx := NewVec3Index()
	materials := make(map[HullKey]string)

	// Prepare all points:
	addFace := func(f *Face) {
		a := idx.Add(f.Vertices[0].Position)
		b := idx.Add(f.Vertices[1].Position)
		c := idx.Add(f.Vertices[2].Position)

		// Now we can add the material to database
		k := HullKey{a, b, c}
		materials[k] = f.MaterialId
	}

	for _, f := range m.Faces {
		addFace(f)
	}

	if len(idx.Data) < 4 {
		return NewModel()
	}

	hull, err := MakeHullFromPoints(idx.Data)
	if err != nil {
		return NewModel()
	}

	res := NewModel()

	// TODO: old materials
	for _, f := range hull {
		res.AddFace(NewFace(idx.Data[f.A], idx.Data[f.B], idx.Data[f.C]))
	}

	return res
}

// Process Minkowski sum for 2 models.
func (m *Model) Minkowski(other *Model) *Model {
	// This algorithm is very slow, but who cares?
	translated := func(v Vec3) *Model {
		res := other.Clone()
		res.Move(v)
		return res
	}

	faceToModel := func(f *Face) *Model {
		var m = translated(f.Vertices[0].Position).Hull(translated(f.Vertices[1].Position))
		return m.Hull(translated(f.Vertices[2].Position))
	}

	res := m.Clone()
	for _, f := range m.Faces {
		res = res.Hull(faceToModel(f))
	}

	return res
}

// Convert model to Csg
func (m *Model) ToCsg() *Csg {
	polygons := make([]*Polygon, 0)

	for _, f := range m.Faces {
		vcs := make([]*Vertex, 0)

		for i := 0; i < 3; i++ {
			vcs = append(vcs, &f.Vertices[i])
		}

		polygons = append(polygons, NewPolygon(vcs, nil, f.MaterialId))
	}

	return NewCsg(polygons...)
}

// Convert Csg to model
func NewModelFromCsg(csg *Csg) *Model {
	// Now create the model and add faces:
	res := NewModel()
	for _, polygon := range csg.Polygons {
		res.addCsgPolygon(polygon, polygon.MaterialId)
	}

	return res
}

func (m *Model) addCsgPolygon(polygon *Polygon, materialId string) error {
	points := make([]Vec3, 0, len(polygon.Vertices))
	for _, v := range polygon.Vertices {
		points = append(points, v.Position)
	}

	triangles, err := TriangulatePolygon3D(points)
	if err != nil {
		return err
	}

	for _, t := range triangles {
		f := NewFaceWithNormals(
			points[t.A], polygon.Vertices[t.A].Normal,
			points[t.B], polygon.Vertices[t.B].Normal,
			points[t.C], polygon.Vertices[t.C].Normal)

		f.Vertices[0].UV = polygon.Vertices[t.A].UV
		f.Vertices[1].UV = polygon.Vertices[t.B].UV
		f.Vertices[2].UV = polygon.Vertices[t.C].UV

		f.MaterialId = materialId

		m.AddFace(f)
	}

	return nil
}
