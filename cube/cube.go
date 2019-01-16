// Package cube implements a Rubik's cube and any other cube e.g. 4x4x4 or 7x7x7.
package cube

import (
	"errors"
	"strings"
)

// Side represents a square of stickers on each side of the cube.
type Side [][]Face

// Cube contains six sides and size n.
type Cube struct {
	N     int
	Sides [6]Side
}

// String implements Stringer interface.
func (c Cube) String() string {
	n := c.N
	tab := strings.Repeat("  ", n) + " "
	str := strings.Builder{}
	str.WriteRune('\n')
	for i := 0; i < n; i++ {
		str.WriteString(tab)
		writeSideLine(&str, &c.Sides[U], i)
		str.WriteRune('\n')
	}
	for i := 0; i < n; i++ {
		writeSideLine(&str, &c.Sides[L], i)
		str.WriteRune(' ')
		writeSideLine(&str, &c.Sides[F], i)
		str.WriteRune(' ')
		writeSideLine(&str, &c.Sides[R], i)
		str.WriteRune(' ')
		writeSideLine(&str, &c.Sides[B], i)
		str.WriteRune('\n')
	}
	for i := 0; i < n; i++ {
		str.WriteString(tab)
		writeSideLine(&str, &c.Sides[D], i)
		str.WriteRune('\n')
	}
	str.WriteRune('\n')
	return str.String()
}

func writeSideLine(str *strings.Builder, s *Side, i int) {
	n := len(*s)
	for j := 0; j < n; j++ {
		str.WriteString((*s)[i][j].String() + " ")
	}
}

// NewCube generates a new cube in the solved state.
func NewCube(n int) (*Cube, error) {
	if n < 1 {
		return nil, errors.New("invalid cube size should be > 0")
	}

	c := &Cube{N: n}
	for i := range c.Sides {
		c.Sides[i] = make([][]Face, n, n)
		for j := range c.Sides[i] {
			c.Sides[i][j] = make([]Face, n, n)
			for k := range c.Sides[i][j] {
				c.Sides[i][j][k] = Face(i)
			}
		}
	}
	return c, nil
}
