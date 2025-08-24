package csgo

import (
	"slices"
)

type Polygon struct {
	Vertices   []*Vertex
	Plane      *Plane
	MaterialId string
}

func NewPolygon(vertices []*Vertex, plane *Plane, materialId ...string) *Polygon {
	mId := ""
	if len(materialId) > 0 {
		mId = materialId[0]
	}

	if plane == nil {
		plane = NewPlaneFromPoints(vertices[0].Position, vertices[1].Position, vertices[2].Position)
	}

	return &Polygon{vertices, plane, mId}
}

func (p *Polygon) Clone() *Polygon {
	vertices := make([]*Vertex, len(p.Vertices))
	for i, v := range p.Vertices {
		vertices[i] = v.Clone()
	}

	return NewPolygon(vertices, nil, p.MaterialId)
}

func (p *Polygon) Flip() {
	for _, v := range p.Vertices {
		v.Flip(p.Plane)
	}
	slices.Reverse(p.Vertices)
	p.Plane.Flip()
}
