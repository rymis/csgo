package csgo

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
)

const floatingPointRx = `([-+0-9.eE]+)`
const vectorRx = floatingPointRx + `\s+` + floatingPointRx + `\s` + floatingPointRx
var facetRx = regexp.MustCompile(`^\s*facet\s+normal\s+` + vectorRx + `\s+outer\s+loop\s+vertex\s+` + vectorRx + `\s+vertex\s+` + vectorRx + `\s+vertex\s+` + vectorRx + `\s+endloop\s+endfacet`)
var solidRx = regexp.MustCompile(`^\s*solid\s+([_a-zA-Z][_a-zA-Z0-9]*)`)
var endsolidRx = regexp.MustCompile(`^\s*endsolid`)

// Save model to STL file
func SaveStl(w io.Writer, model *Model) error {
	// TODO: error handling
	fmt.Fprintln(w, "solid ScadSharp")
	for _, f := range model.Faces {
		n := f.Normal()
		fmt.Fprintf(w, "  facet normal %f %f %f\n", n.X(), n.Y(), n.Z())
		fmt.Fprintln(w, "    outer loop")
		fmt.Fprintf(w, "       vertex %f %f %f\n", f.Vertices[0].Position.X(), f.Vertices[0].Position.Y(), f.Vertices[0].Position.Z())
		fmt.Fprintf(w, "       vertex %f %f %f\n", f.Vertices[1].Position.X(), f.Vertices[1].Position.Y(), f.Vertices[1].Position.Z())
		fmt.Fprintf(w, "       vertex %f %f %f\n", f.Vertices[2].Position.X(), f.Vertices[2].Position.Y(), f.Vertices[2].Position.Z())
		fmt.Fprintln(w, "    endloop")
		fmt.Fprintln(w, "  endfacet")
	}
	fmt.Fprintln(w, "endsolid")

	return nil
}

// Save binary STL model to writer
func SaveStlBin(w io.Writer, m *Model) error {
	// First write header:
	header := make([]byte, 80)
	text := []byte("ScadView Model\n")
	copy(header[:len(text)], text)

	_, err := w.Write(header) // 80 bytes header
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Faces)))
	if err != nil {
		return err
	}

	writeVec := func (v Vec3) error {
		for i := 0; i < 3; i++ {
			err := binary.Write(w, binary.LittleEndian, v[0])
			if err != nil {
				return err
			}
		}

		return nil
	};

	// Triangles:
	for _, f := range m.Faces {
		n := f.Normal()

		err = writeVec(n)
		if err != nil {
			return err
		}
		err = writeVec(f.Vertices[0].Position)
		if err != nil {
			return err
		}
		err = writeVec(f.Vertices[1].Position)
		if err != nil {
			return err
		}
		err = writeVec(f.Vertices[2].Position)
		if err != nil {
			return err
		}
		err = binary.Write(w, binary.LittleEndian, uint16(0))
		if err != nil {
			return err
		}
		// TODO:
		// The VisCAM and SolidView software packages use the two "attribute byte count"
		// bytes at the end of every triangle to store a 15-bit RGB color:
		//   bits 0–4 are the intensity level for blue (0–31),
		//   bits 5–9 are the intensity level for green (0–31),
		//   bits 10–14 are the intensity level for red (0–31),
		//   bit 15 is 1 if the color is valid, or 0 if the color is not valid (as with normal STL files).
	}

	return nil
}

// Load model from STL textual format
func LoadStlText(r io.Reader) (*Model, error) {
	var x [12]float32
	res := NewModel()
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	indices := solidRx.FindSubmatchIndex(data)
	if indices == nil {
		return nil, fmt.Errorf("Invalid model format")
	}

	name := string(data[indices[2]:indices[3]])
	if name == "facet" { // No name present
		data = data[indices[2]:]
	} else {
		data = data[indices[3]:]
	}

	// Load facets:
	for {
		if endsolidRx.Match(data) {
			break
		}

		indices = facetRx.FindSubmatchIndex(data)
		if indices == nil {
			return nil, fmt.Errorf("Invalid STL model file format")
		}

		if len(indices) != 26 {
			return nil, fmt.Errorf("Invalid STL model file format: %d", len(indices))
		}

		for i := 0; i < 12; i++ {
			l := indices[i * 2 + 2]
			r := indices[i * 2 + 3]
			if l < 0 || r < 0 {
				return nil, fmt.Errorf("Invalid STL model file format")
			}

			val := data[l: r]
			xx, err := strconv.ParseFloat(string(val), 32)
			if err != nil {
				return nil, err
			}
			x[i] = float32(xx)
		}

		f := NewFace(Vec3{x[3], x[4], x[5]}, Vec3{x[6], x[7], x[8]}, Vec3{x[9], x[10], x[11]})
		res.AddFace(f)

		data = data[indices[1]:]
	}

	return res, nil
}
