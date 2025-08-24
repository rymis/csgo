package csgo

import (
	"fmt"
)

// TODO: we do not need to have Face struct here.

// Idea of this implementation came from:
// https://cw.fel.cvut.cz/wiki/_media/misc/projects/oppa_oi_english/courses/ae4m39vg/lectures/05-convexhull-3d.pdf

// Triangle contructed in Hull
type HullKey struct {
	A, B, C int
}

// Create a HullKey by 3 indices. Hull key index is sorted, so the first integer is the minimal one.
func NewHullKey(a, b, c int) HullKey {
	// Put the lowest index first
	if a < b && a < c {
		return HullKey{a, b, c}
	}

	if b < a && b < c {
		return HullKey{b, c, a}
	}

	return HullKey{c, a, b}
}

const (
	hullFront    uint = 0
	hullBack     uint = 1
	hullCoplanar uint = 2
)

type hullState struct {
	Points []Vec3
	Faces  map[HullKey]*hullFace
}

type hullFace struct {
	Hull    *hullState
	A, B, C int
}

func (hf *hullFace) GetA() Vec3 {
	return hf.Hull.Points[hf.A]
}

func (hf *hullFace) GetB() Vec3 {
	return hf.Hull.Points[hf.B]
}

func (hf *hullFace) GetC() Vec3 {
	return hf.Hull.Points[hf.C]
}

func newHullFace(h *hullState, a, b, c int) *hullFace {
	res := &hullFace{}
	res.Hull = h

	// Put the lowest index first
	if a < b && a < c {
		res.A = a
		res.B = b
		res.C = c
	} else if b < a && b < c {
		res.A = b
		res.B = c
		res.C = a
	} else {
		res.A = c
		res.B = a
		res.C = b
	}

	return res
}

func (hf *hullFace) Normal() Vec3 {
	return hf.GetC().Sub(hf.GetA()).Cross(hf.GetB().Sub(hf.GetA())).Normalize()
}

func (hf *hullFace) Dist(x Vec3) float32 {
	return hf.Normal().Dot(x.Sub(hf.GetA()))
}

func (hf *hullFace) PointClass(x Vec3) uint {
	v := hf.Dist(x)
	if v > Epsilon {
		return hullFront
	}

	if v < -Epsilon {
		return hullBack
	}

	return hullCoplanar
}

func (hf *hullFace) Key() HullKey {
	return HullKey{hf.A, hf.B, hf.C}
}

func newHullState(points []Vec3) *hullState {
	return &hullState{
		points,
		make(map[HullKey]*hullFace),
	}
}

func (s *hullState) IsInside(p Vec3) bool {
	for _, f := range s.Faces {
		if f.PointClass(p) == hullBack {
			return false
		}
	}

	return true
}

func (s *hullState) LineDistance(a, b, c int) float32 {
	p := (s.Points[b].X()-s.Points[a].X())*(s.Points[a].Y()-s.Points[c].Y()) - (s.Points[a].X()-s.Points[c].X())*(s.Points[b].Y()-s.Points[a].Y())
	q := s.Points[b].Sub(s.Points[a]).LenSqr()

	return p * p / q
}

func (s *hullState) AddFace(a, b, c int) {
	f := newHullFace(s, a, b, c)
	s.Faces[f.Key()] = f
}

func (s *hullState) InitialHull() error {
	splitA := 0
	splitB := 1

	// Find the most distant points:
	{
		maxDist := s.Points[splitA].Sub(s.Points[splitB]).LenSqr()
		for i := splitA; i < len(s.Points); i++ {
			for j := i + 1; j < len(s.Points); j++ {
				d := s.Points[i].Sub(s.Points[j]).LenSqr()
				if d > maxDist {
					maxDist = d
					splitA = i
					splitB = j
				}
			}
		}
	}

	splitC := 0
	for splitC == splitA || splitC == splitB {
		splitC++
	}

	{
		maxDist := s.LineDistance(splitA, splitB, splitC)
		for i := 0; i < len(s.Points); i++ {
			if i == splitA || i == splitB {
				continue
			}

			d := s.LineDistance(splitA, splitB, i)
			if d > maxDist {
				maxDist = d
				splitC = i
			}
		}
	}

	// OK: we have a plane here
	a := s.Points[splitA]
	n := s.Points[splitC].Sub(a).Cross(s.Points[splitB].Sub(a)).Normalize()
	planeDist := func(i int) float32 {
		return n.Dot(s.Points[i].Sub(a))
	}

	// Add front and back points to the plane:
	{
		front := 0
		back := 0
		minDist := planeDist(0)
		maxDist := minDist

		for i := 1; i < len(s.Points); i++ {
			d := planeDist(i)
			if d > maxDist {
				maxDist = d
				front = i
			}

			if d < minDist {
				minDist = d
				back = i
			}
		}

		if maxDist > 0.0 {
			s.AddFace(splitA, splitB, front)
			s.AddFace(splitB, splitC, front)
			s.AddFace(splitC, splitA, front)
		} else {
			s.AddFace(splitA, splitB, splitC)
		}

		if minDist < 0.0 {
			s.AddFace(splitB, splitA, back)
			s.AddFace(splitC, splitB, back)
			s.AddFace(splitA, splitC, back)
		} else {
			s.AddFace(splitC, splitB, splitA)
		}
	}

	if len(s.Faces) < 3 {
		return fmt.Errorf("Trying to make degraded Hull")
	}

	return nil
}

func (s *hullState) AddEyePoint(p int) {
	type sideInfo struct {
		A, B int
	}
	sides := make(map[sideInfo]int)
	toRemove := make(map[HullKey]bool)

	getSide := func(a, b int) int {
		v, ok := sides[sideInfo{a, b}]
		if !ok {
			return 0
		}
		return v
	}

	for vk, vf := range s.Faces {
		cls := vf.PointClass(s.Points[p])
		if cls == hullFront || cls == hullCoplanar {
			sides[sideInfo{vf.A, vf.B}] = getSide(vf.A, vf.B) + 1
			sides[sideInfo{vf.B, vf.C}] = getSide(vf.B, vf.C) + 1
			sides[sideInfo{vf.C, vf.A}] = getSide(vf.C, vf.A) + 1

			toRemove[vk] = true
		}
	}

	// Erase invisible faces:
	for k := range toRemove {
		delete(s.Faces, k)
	}

	// Add new faces:
	for sk, sv := range sides {
		if sv == 1 {
			_, ok := sides[sideInfo{sk.B, sk.A}]
			if !ok {
				s.AddFace(sk.A, sk.B, p)
			}
		}
	}
}

func (s *hullState) Create() ([]HullKey, error) {
	if len(s.Points) < 4 {
		return nil, fmt.Errorf("Trying to make degraded hull (less than 4 points)")
	}

	err := s.InitialHull()
	if err != nil {
		return nil, err
	}

	used := make([]bool, len(s.Points))
	for _, f := range s.Faces {
		used[f.A] = true
		used[f.B] = true
		used[f.C] = true
	}

	for i := 0; i < len(s.Points); i++ {
		if used[i] {
			continue
		}

		if s.IsInside(s.Points[i]) {
			continue
		}

		s.AddEyePoint(i)
	}

	res := make([]HullKey, 0)
	for _, f := range s.Faces {
		res = append(res, HullKey{f.A, f.B, f.C})
	}

	return res, nil
}

// Generate convex hull by list of 3D points.
func MakeHullFromPoints(points []Vec3) ([]HullKey, error) {
	impl := newHullState(points)

	return impl.Create()
}
