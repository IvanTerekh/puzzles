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
			face: R,
			dir:  Cw,
		},
		Cube:     *cube3,
		valid:    true,
		rotation: false,
	},
	{
		Move: Move{
			N:    1,
			face: F,
			dir:  Acw,
		},
		Cube:     *cube3,
		valid:    true,
		rotation: false,
	},
	{
		Move: Move{
			N:    2,
			face: B,
			dir:  Two,
		},
		Cube:     *cube3,
		valid:    true,
		rotation: false,
	},
	{
		Move: Move{
			N:    4,
			face: R,
			dir:  Cw,
		},
		Cube:     *cube3,
		valid:    false,
		rotation: false,
	},
	{
		Move: Move{
			N:    4,
			face: R,
			dir:  Cw,
		},
		Cube:     *cube7,
		valid:    true,
		rotation: false,
	},
	{
		Move: Move{
			N:    7,
			face: R,
			dir:  Cw,
		},
		Cube:     *cube7,
		valid:    true,
		rotation: true,
	},
	{
		Move: Move{
			N:    7,
			face: R,
			dir:  Cw,
		},
		Cube:     *cube3,
		valid:    false,
		rotation: false,
	},
	{
		Move: Move{
			N:    3,
			face: L,
			dir:  Cw,
		},
		Cube:     *cube3,
		valid:    true,
		rotation: true,
	},
}
