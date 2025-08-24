package csgo

// This file is based on csg.js written by Eval Wallace. Original module is licensed under the MIT license:
//   Copyright (c) 2011 Evan Wallace (http://madebyevan.com/), under the MIT license.
// I think it is fare to say that this file is authored by me:
//   Copyright (c) 2025 Mikhail Ryzhov (rymiser@gmail.com), under the MIT license.

import (
	"math/rand"
)

// Holds a binary space partition tree representing a 3D solid. Two solids can
// be combined using the Union(), Subtract(), and Intersect() methods.
type Csg struct {
	Polygons []*Polygon
}

// Create new Csg object from polygons.
func NewCsg(polygons ...*Polygon) *Csg {
	return &Csg{polygons}
}

// Create a full copy of the Csg
func (csg *Csg) Clone() *Csg {
	pols := make([]*Polygon, 0)
	for _, pol := range csg.Polygons {
		pols = append(pols, pol.Clone())
	}

	return NewCsg(pols...)
}

/*
	Return a new CSG solid representing space in either this solid or in the
	solid `csg`. Neither this solid nor the solid `csg` are modified.

	A.Union(B)

	    +-------+            +-------+
	    |       |            |       |
	    |   A   |            |       |
	    |    +--+----+   =   |       +----+
	    +----+--+    |       +----+       |
	         |   B   |            |       |
	         |       |            |       |
	         +-------+            +-------+
*/
func (csg *Csg) Union(other *Csg) *Csg {
	a := newCsgNode(csg.Clone().Polygons...)
	b := newCsgNode(other.Clone().Polygons...)
	a.ClipTo(b)
	b.ClipTo(a)
	b.Invert()
	b.ClipTo(a)
	b.Invert()
	a.Build(b.AllPolygons(), 0)

	return NewCsg(a.AllPolygons()...)
}

/*
Return a new CSG solid representing space in this solid but not in the
solid `csg`. Neither this solid nor the solid `csg` are modified.

	A.Subtract(B)

	+-------+            +-------+
	|       |            |       |
	|   A   |            |       |
	|    +--+----+   =   |    +--+
	+----+--+    |       +----+
	     |   B   |
	     |       |
	     +-------+
*/
func (csg *Csg) Subtract(other *Csg) *Csg {
	a := newCsgNode(csg.Clone().Polygons...)
	b := newCsgNode(other.Clone().Polygons...)
	a.Invert()
	a.ClipTo(b)
	b.ClipTo(a)
	b.Invert()
	b.ClipTo(a)
	b.Invert()
	a.Build(b.AllPolygons(), 0)
	a.Invert()

	return NewCsg(a.AllPolygons()...)
}

/*
Return a new CSG solid representing space both this solid and in the
solid `csg`. Neither this solid nor the solid `csg` are modified.

	A.Intersect(B)

	+-------+
	|       |
	|   A   |
	|    +--+----+   =   +--+
	+----+--+    |       +--+
	     |   B   |
	     |       |
	     +-------+
*/
func (csg *Csg) Intersect(other *Csg) *Csg {
	a := newCsgNode(csg.Clone().Polygons...)
	b := newCsgNode(other.Clone().Polygons...)
	a.Invert()
	b.ClipTo(a)
	b.Invert()
	a.ClipTo(b)
	b.ClipTo(a)
	a.Build(b.AllPolygons(), 0)
	a.Invert()

	return NewCsg(a.AllPolygons()...)
}

// Maximum build level to prevent infinite recursion
const MaxBuildLevel = 65536

type csgNode struct {
	Polygons    []*Polygon
	Plane       *Plane
	front, back *csgNode
}

func newCsgNode(polygons ...*Polygon) *csgNode {
	res := &csgNode{make([]*Polygon, 0), nil, nil, nil}
	res.Build(polygons, 0)
	return res
}

func (n *csgNode) Build(polygons []*Polygon, level int) {
	if level > MaxBuildLevel {
		println("WARNING: Build level is too high")
		return
	}

	if len(polygons) == 0 {
		return
	}

	if n.Plane == nil {
		rnd := rand.Intn(len(polygons))
		n.Plane = polygons[rnd].Plane.Clone()
	}

	front := make([]*Polygon, 0)
	back := make([]*Polygon, 0)

	for _, polygon := range polygons {
		n.Plane.SplitPolygon(polygon,
			func(p *Polygon) { n.Polygons = append(n.Polygons, p) },
			func(p *Polygon) { n.Polygons = append(n.Polygons, p) },
			func(p *Polygon) { front = append(front, p) },
			func(p *Polygon) { back = append(back, p) })
	}

	if len(front) != 0 {
		if n.front == nil {
			n.front = newCsgNode()
		}
		n.front.Build(front, level+1)
	}

	if len(back) != 0 {
		if n.back == nil {
			n.back = newCsgNode()
		}
		n.back.Build(back, level+1)
	}
}

func (n *csgNode) AllPolygons() []*Polygon {
	var res = make([]*Polygon, 0)
	res = append(res, n.Polygons...)
	if n.back != nil {
		res = append(res, n.back.AllPolygons()...)
	}
	if n.front != nil {
		res = append(res, n.front.AllPolygons()...)
	}
	return res
}

func (n *csgNode) Clone() *csgNode {
	res := newCsgNode()

	if n.Plane != nil {
		res.Plane = n.Plane.Clone()
	}

	if n.front != nil {
		res.front = n.front.Clone()
	}

	if n.back != nil {
		res.back = n.back.Clone()
	}

	for _, polygon := range n.Polygons {
		res.Polygons = append(res.Polygons, polygon.Clone())
	}

	return res
}

// Convert solid space to empty space and empty space to solid space.
func (n *csgNode) Invert() {
	for _, p := range n.Polygons {
		p.Flip()
	}

	if n.Plane != nil {
		n.Plane.Flip()
	}

	if n.front != nil {
		n.front.Invert()
	}

	if n.back != nil {
		n.back.Invert()
	}

	tmp := n.back
	n.back = n.front
	n.front = tmp
}

// Recursively remove all polygons in `polygons` that are inside this BSP tree
func (n *csgNode) ClipPolygons(polygons []*Polygon) []*Polygon {
	if n.Plane == nil {
		return append(make([]*Polygon, 0, len(polygons)), polygons...)
	}

	front := make([]*Polygon, 0)
	back := make([]*Polygon, 0)

	addFront := func(p *Polygon) { front = append(front, p) }
	addBack := func(p *Polygon) { back = append(back, p) }

	for _, p := range polygons {
		n.Plane.SplitPolygon(p, addFront, addBack, addFront, addBack)
	}

	if n.front != nil {
		front = n.front.ClipPolygons(front)
	}

	if n.back != nil {
		back = n.back.ClipPolygons(back)
		front = append(front, back...)
	}

	return front
}

// Remove all polygons in this BSP tree that are inside the other BSP tree `bsp`.
func (n *csgNode) ClipTo(bsp *csgNode) {
	n.Polygons = bsp.ClipPolygons(n.Polygons)

	if n.front != nil {
		n.front.ClipTo(bsp)
	}

	if n.back != nil {
		n.back.ClipTo(bsp)
	}
}
