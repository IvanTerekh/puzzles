package cube

const (
	// Cw is clockwise direction
	Cw = iota
	// Acw is anticlockwise direction
	Acw
	// Two is double turn
	Two
)

// Dir represents move direction. Can by clockwise, anticlockwise and double turn.
type Dir byte

// String is Stringer interface implementation.
func (d Dir) String() string {
	switch d {
	case Cw:
		return "Clockwise"
	case Acw:
		return "Anticlockwise"
	case Two:
		return "Double turn"
	default:
		return "Invalid direction"
	}
}
