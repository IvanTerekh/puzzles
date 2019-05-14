package cube

// Move represents a move of n layers on a cube with given direction. Move of all layers is meant to be a rotation.
type Move struct {
	N int
	face
	dir
}

// validFor checks if m is a valid move for cube.
func (m *Move) validFor(c Cube) bool {
	return m.N > 0 && m.N <= c.n
}

// rotationFor check if m is rotation for a cube.
func (m *Move) rotationFor(c Cube) bool {
	return m.N == c.n
}
