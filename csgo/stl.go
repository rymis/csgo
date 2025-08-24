package csgo

import (
	"fmt"
	"io"
)

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
/*
public static void SaveBin(Model m, System.IO.BinaryWriter w)
{
	// First write header:
	var header = new byte[80];
	string text = "ScadView Model\n";
	for (int i = 0; i < text.Length; ++i) {
		header[i] = (byte)text[i];
	}

	w.Write(header);          // 80 bytes header
	w.Write((uint)m.Count()); // number of triangles
	var writeVec = (Linalg.Vec3 v) => {
		w.Write(v.X);
		w.Write(v.Y);
		w.Write(v.Z);
	};

	// Triangles:
	for (int i = 0; i < m.Count(); ++i) {
		var f = m.GetFace(i);
		var n = f.Normal();
		writeVec(n);
		writeVec(f.Vertices[0]);
		writeVec(f.Vertices[1]);
		writeVec(f.Vertices[2]);
		w.Write((short)0);
		// TODO:
		// The VisCAM and SolidView software packages use the two "attribute byte count"
		// bytes at the end of every triangle to store a 15-bit RGB color:
		//   bits 0–4 are the intensity level for blue (0–31),
		//   bits 5–9 are the intensity level for green (0–31),
		//   bits 10–14 are the intensity level for red (0–31),
		//   bit 15 is 1 if the color is valid, or 0 if the color is not valid (as with normal STL files).
	}
}
*/
