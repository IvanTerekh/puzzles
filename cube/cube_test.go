package cube

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCube(t *testing.T) {
	n := 3
	cube, err := NewCube(n)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	fmt.Println(cube)

	l := len(cube.Sides[F])
	if l != n {
		t.Fatalf("Wrong side size allocated. got: %d, expected: %d", l, n)
	}
	for i := 0; i < n; i++ {
		l = len(cube.Sides[F][i])
		if l != n {
			t.Fatalf("Wrong side size allocated. got: %d, expected: %d", l, n)
		}
	}

	frontSide := side([][]face{{F, F, F}, {F, F, F}, {F, F, F}})
	if !reflect.DeepEqual(cube.Sides[F], frontSide) {
		t.Fatalf("Front side is not properly assigned, expected: %v, got: %v", frontSide, cube.Sides[F])
	}

	downSide := side([][]face{{D, D, D}, {D, D, D}, {D, D, D}})
	if !reflect.DeepEqual(cube.Sides[D], downSide) {
		t.Fatalf("Down side is not properly assigned, expected: %v, got: %v", downSide, cube.Sides[D])
	}

	leftSide := side([][]face{{L, L, L}, {L, L, L}, {L, L, L}})
	if !reflect.DeepEqual(cube.Sides[L], leftSide) {
		t.Fatalf("Left side is not properly assigned, expected: %v, got: %v", leftSide, cube.Sides[L])
	}

	_, err = NewCube(0)
	if err == nil {
		t.Fatalf("Expected error while generating zero-layer cube, got nil.")
	}

	_, err = NewCube(-6)
	if err == nil {
		t.Fatalf("Expected error while generating cube -6 by -6 by -6, got nil.")
	}
}

func TestValid(t *testing.T) {
	for _, move := range movesForValidation {
		if move.validFor(move.Cube) != move.valid {
			t.Fatalf("Wrong validation result for %v,\n expexted: %t, got: %t.",
				move, move.valid, move.validFor(move.Cube))
		}
	}
}

func TestRotation(t *testing.T) {
	for _, move := range movesForValidation {
		if move.rotationFor(move.Cube) != move.rotation {
			t.Fatalf("Wrong validation result for %v,\n expexted: %t, got: %t.",
				move, move.rotation, move.rotationFor(move.Cube))
		}
	}
}
