package cube

// Move represents a move of N layers on a cube with given direction. Move of all layers is meant to be a rotation.
type Move struct {
	N int
	Face
	Dir
}

// ValidFor checks if m is a valid move for cube.
func (m *Move) ValidFor(c Cube) bool {
	return m.N > 0 && m.N <= c.N
}

// RotationFor check if m is rotation for a cube.
func (m *Move) RotationFor(c Cube) bool {
	return m.N == c.N
}
