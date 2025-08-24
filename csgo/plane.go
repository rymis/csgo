package csgo

// 3D Plane representation.
type Plane struct {
	Normal Vec3
	W      float32
}

// Create a new place by normal and distance from the (0, 0, 0).
func NewPlane(normal Vec3, w float32) *Plane {
	return &Plane{normal, w}
}

// Create new plane by 3 points.
func NewPlaneFromPoints(a, b, c Vec3) *Plane {
	n := b.Sub(a).Cross(c.Sub(a)).Normalize()
	return NewPlane(n, n.Dot(a))
}

// Create a full copy of this plane
func (p *Plane) Clone() *Plane {
	return NewPlane(p.Normal, p.W)
}

// Flip the plane.
func (p *Plane) Flip() {
	p.Normal = p.Normal.Mul(-1.0)
	p.W = -p.W
}

// Mirror the plane by some vector.
func (p *Plane) Mirror(v Vec3) Vec3 {
	n := 2.0 * p.Normal.Dot(v)
	return v.Sub(p.Normal.Mul(n))
}

// Split polygon by the plane. Function adds polygon to the correct set. This function is used in CSG algorithms.
func (p *Plane) SplitPolygon(polygon *Polygon, addCoplanarFront, addCoplanarBack, addFront, addBack func(p *Polygon)) {
	const Coplanar uint32 = 0
	const Back uint32 = 1
	const Front uint32 = 2
	const Spanning uint32 = 3

	// Classify polygon points and polygon:
	polygonType := uint32(0)
	types := make([]uint32, 0)

	for _, vertex := range polygon.Vertices {
		t := p.Normal.Dot(vertex.Position) - p.W
		pType := uint32(0)

		if t < -Epsilon {
			pType = Back
		} else if t > Epsilon {
			pType = Front
		} else {
			pType = Coplanar
		}

		polygonType |= pType
		types = append(types, pType)
	}

	// Add polygon to correct list:
	if polygonType == Coplanar {
		if p.Normal.Dot(polygon.Plane.Normal) > 0.0 {
			addCoplanarFront(polygon)
		} else {
			addCoplanarBack(polygon)
		}
	} else if polygonType == Front {
		addFront(polygon)
	} else if polygonType == Back {
		addBack(polygon)
	} else {
		// Spanning polygon
		f := make([]*Vertex, 0)
		b := make([]*Vertex, 0)

		for i := range polygon.Vertices {
			j := (i + 1) % len(polygon.Vertices)
			ti := types[i]
			tj := types[j]
			vi := polygon.Vertices[i]
			vj := polygon.Vertices[j]

			if ti != Back {
				f = append(f, vi)
			}

			if ti != Front {
				if ti != Back {
					b = append(b, vi.Clone())
				} else {
					b = append(b, vi)
				}
			}

			if (ti | tj) == Spanning {
				t := (p.W - p.Normal.Dot(vi.Position)) / p.Normal.Dot(vj.Position.Sub(vi.Position))
				v := vi.Interpolate(vj, t)
				f = append(f, v)
				b = append(b, v.Clone())
			}
		}

		if len(f) >= 3 {
			addFront(NewPolygon(f, nil, polygon.MaterialId))
		}

		if len(b) >= 3 {
			addBack(NewPolygon(b, nil, polygon.MaterialId))
		}
	}
}
