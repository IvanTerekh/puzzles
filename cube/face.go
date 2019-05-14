package cube

const (
	// F is Front.
	F face = iota
	// R is right.
	R
	// B is back.
	B
	// L is left.
	L
	// U is up.
	U
	// D is down.
	D
)

// face represents a color on a cube. As cube's color can vary instead of using particular colors
// sides colors used, e.g. front side color, right side color etc. Possible values are F, R, B, L, U, D.
type face byte

// String is an implementation of Stringer interface.
func (f face) String() string {
	switch f {
	case F:
		return "F"
	case R:
		return "R"
	case B:
		return "B"
	case L:
		return "L"
	case U:
		return "U"
	case D:
		return "D"
	default:
		return "Invalid Move"
	}
}

func (f face) opposite() face {
	opposites := [6]face{D, R, B, L, F, U}
	return opposites[f]
}
