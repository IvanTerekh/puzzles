package cube

import (
	"github.com/getlantern/deepcopy"
	"reflect"
	"testing"
)

var startPosition = side([][]face{
	{D, L, L, B, F, D, R},
	{L, D, B, L, L, L, R},
	{R, L, F, F, U, R, F},
	{D, D, L, U, D, B, L},
	{D, U, R, B, D, R, U},
	{U, D, F, F, U, F, B},
	{L, L, B, R, D, L, L},
})

var turns = []struct {
	dir
	side
}{
	{
		dir: Cw,
		side: side([][]face{
			{L, U, D, D, R, L, D},
			{L, D, U, D, L, D, L},
			{B, F, R, L, F, B, L},
			{R, F, B, U, F, L, B},
			{D, U, D, D, U, L, F},
			{L, F, R, B, R, L, D},
			{L, B, U, L, F, R, R},
		}),
	},
	{
		dir: Acw,
		side: side([][]face{
			{R, R, F, L, U, B, L},
			{D, L, R, B, R, F, L},
			{F, L, U, D, D, U, D},
			{B, L, F, U, B, F, R},
			{L, B, F, L, R, F, B},
			{L, D, L, D, U, D, L},
			{D, L, R, D, D, U, L},
		}),
	},
}

func TestTurn(t *testing.T) {
	var side side
	for _, testcase := range turns {
		err := deepcopy.Copy(&side, startPosition)
		if err != nil {
			panic(err)
		}
		side.turn(testcase.dir)
		if !reflect.DeepEqual(side, testcase.side) {
			t.Errorf("wrong turn result:\n%v\nexpected:\n%v", side, testcase.side)
		}
	}
}
