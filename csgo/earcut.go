package csgo

import (
	"fmt"
)

// Simple implementation of Ear-Cut algorithm.

// Triangle as 3 indices
type EarcutTriangle struct {
	A, B, C int
}

// Check if point x is in triangle ABC
func IsInTriangle(x, a, b, c Vec2) bool {
	side := func(f float32) uint {
		if f < -Epsilon {
			return 1
		} else if f > Epsilon {
			return 2
		} else {
			return 3
		}
	}

	cross := func(a, b Vec2) float32 {
		return a.X()*b.Y() - a.Y()*b.X()
	}

	aSide := side(cross(b.Sub(a), x.Sub(a)))
	bSide := side(cross(c.Sub(b), x.Sub(b)))
	cSide := side(cross(a.Sub(c), x.Sub(c)))

	return (aSide & bSide & cSide) != 0
}

// Function that triangulates polygon.
func TriangulatePolygon(polygon []Vec2) ([]EarcutTriangle, error) {
	type earIndex struct {
		V Vec2
		I int
	}

	if len(polygon) < 3 {
		return nil, fmt.Errorf("Need at least 3 points to cut ears")
	}

	res := make([]EarcutTriangle, 0, len(polygon)-2)
	idx := make([]earIndex, 0, len(polygon))

	for i, p := range polygon {
		idx = append(idx, earIndex{p, i})
	}

	for len(idx) > 3 {
		oldSize := len(idx)

		for i := 0; len(idx) > 3 && i < len(idx); i++ {
			a := idx[i].V
			b := idx[(i+1)%len(idx)].V
			c := idx[(i+2)%len(idx)].V

			isEar := true
			for j := 3; j < len(idx); j++ {
				if IsInTriangle(idx[(i+j)%len(idx)].V, a, b, c) {
					isEar = false
					break
				}
			}

			if isEar {
				t := EarcutTriangle{}

				t.A = idx[i].I
				t.B = idx[(i+1)%len(idx)].I
				t.C = idx[(i+2)%len(idx)].I

				res = append(res, t)

				k := (i + 1) % len(idx)
				idx = append(idx[:k], idx[k + 1:]...)

				break
			}
		}

		if oldSize == len(idx) {
			if len(res) > 0 {
				// We have enough :)
				return res, nil
			}

			return nil, fmt.Errorf("Theorem about ears does not work for your polygone")
		}
	}

	lastT := EarcutTriangle{}
	lastT.A = idx[0].I
	lastT.B = idx[1].I
	lastT.C = idx[2].I

	res = append(res, lastT)

	return res, nil
}

// Triangulate polygon in 3D space
func TriangulatePolygon3D(polygon []Vec3) ([]EarcutTriangle, error) {
	if len(polygon) < 3 {
		return nil, fmt.Errorf("Need at least 3 points to cut ears")
	}

	n := polygon[2].Sub(polygon[0]).Cross(polygon[1].Sub(polygon[0])).Normalize()
	i := polygon[1].Sub(polygon[0]).Normalize()
	j := i.Cross(n)

	points := make([]Vec2, 0)
	for _, p := range polygon {
		points = append(points, Vec2{p.Dot(i), p.Dot(j)})
	}

	return TriangulatePolygon(points)
}
