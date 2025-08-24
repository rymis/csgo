package csgo

type quantizedVec struct {
	X, Y, Z int
}

// Utility class used in model exporters
type Vec3Index struct {
	index map[quantizedVec]int
	Data  []Vec3
}

// Create a new index
func NewVec3Index() *Vec3Index {
	return &Vec3Index{
		make(map[quantizedVec]int),
		make([]Vec3, 0, 64),
	}
}

// Get index of the vector. If vector is not found returns -1
func (v3i *Vec3Index) Index(v Vec3) int {
	idx, ok := v3i.index[quantizeVec(v)]
	if !ok {
		return -1
	}
	return idx
}

// Get or add vector to the index. Returns index of the vector.
func (v3i *Vec3Index) Add(v Vec3) int {
	q := quantizeVec(v)
	idx, ok := v3i.index[q]
	if !ok {
		idx = len(v3i.Data)
		v3i.index[q] = idx
		v3i.Data = append(v3i.Data, v)
	}
	return idx
}

func quantizeVec(v Vec3) quantizedVec {
	x := int(v[0] * 16384.0)
	y := int(v[1] * 16384.0)
	z := int(v[2] * 16384.0)
	return quantizedVec{x, y, z}
}
