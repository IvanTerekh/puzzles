package cube

var cube3, _ = NewCube(3)
var cube7, _ = NewCube(7)

var movesForValidation = []struct {
	Move
	Cube     Cube
	valid    bool
	rotation bool
}{
	{
		Move: Move{
			N:    1,
			Face: R,
			Dir:  Cw,
		},
		Cube:     *cube3,
		valid:    true,
		rotation: false,
	},
	{
		Move: Move{
			N:    1,
			Face: F,
			Dir:  Acw,
		},
		Cube:     *cube3,
		valid:    true,
		rotation: false,
	},
	{
		Move: Move{
			N:    2,
			Face: B,
			Dir:  Two,
		},
		Cube:     *cube3,
		valid:    true,
		rotation: false,
	},
	{
		Move: Move{
			N:    4,
			Face: R,
			Dir:  Cw,
		},
		Cube:     *cube3,
		valid:    false,
		rotation: false,
	},
	{
		Move: Move{
			N:    4,
			Face: R,
			Dir:  Cw,
		},
		Cube:     *cube7,
		valid:    true,
		rotation: false,
	},
	{
		Move: Move{
			N:    7,
			Face: R,
			Dir:  Cw,
		},
		Cube:     *cube7,
		valid:    true,
		rotation: true,
	},
	{
		Move: Move{
			N:    7,
			Face: R,
			Dir:  Cw,
		},
		Cube:     *cube3,
		valid:    false,
		rotation: false,
	},
	{
		Move: Move{
			N:    3,
			Face: L,
			Dir:  Cw,
		},
		Cube:     *cube3,
		valid:    true,
		rotation: true,
	},
}
