package csgo

// Vertex information
type Vertex struct {
	// Position of this vertex.
	Position Vec3
	// Vertex normal
	Normal   Vec3
	// Texture coordinates of this vertex
	UV       Vec2
}

// Create a new vertex
func NewVertex(pos, normal Vec3, uv ...Vec2) *Vertex {
	uv2 := Vec2{0.0, 0.0}
	if len(uv) > 0 {
		uv2 = uv[0]
	}

	return &Vertex{pos, normal, uv2}
}

// Create a copy of the vertex
func (v *Vertex) Clone() *Vertex {
	return NewVertex(v.Position, v.Normal, v.UV)
}

// Flip this vertex over the plane
func (v *Vertex) Flip(p *Plane) {
	v.Normal = p.Mirror(v.Normal)
}

func lerp(x, a Vec3, t float32) Vec3 {
	return x.Add(a.Sub(x).Mul(t))
}

func lerp2(x, a Vec2, t float32) Vec2 {
	return x.Add(a.Sub(x).Mul(t))
}

// Interpolate this vector to get intermediate vertex on the edge.
func (v *Vertex) Interpolate(other *Vertex, t float32) *Vertex {
	return NewVertex(
		lerp(v.Position, other.Position, t),
		lerp(v.Normal, other.Normal, t),
		lerp2(v.UV, other.UV, t))
}
