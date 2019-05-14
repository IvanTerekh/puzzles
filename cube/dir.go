package cube

const (
	// Cw is clockwise direction
	Cw = iota
	// Acw is anticlockwise direction
	Acw
	// Two is double turn
	Two
)

// dir represents move direction. Can by clockwise, anticlockwise and double turn.
type dir byte

// String is Stringer interface implementation.
func (d dir) String() string {
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

func (d dir) opposite() dir {
	switch d {
	case Cw:
		return Acw
	case Acw:
		return Cw
	default:
		return d
	}
}
