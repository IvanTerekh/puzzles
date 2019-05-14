// Package cube implements a Rubik's cube and any other cube e.g. 4x4x4 or 7x7x7.
package cube

import (
	"errors"
	"strings"
)

// Cube contains six sides and size n.
// Side matrices are located the following way:
//              0,0 0,1 0,2
//              1,0 1,1 1,2
//              2,0 2,1 2,2
// 0,0 0,1 0,2  0,0 0,1 0,2  0,0 0,1 0,2  0,0 0,1 0,2
// 1,0 1,1 1,2  1,0 1,1 1,2  1,0 1,1 1,2  1,0 1,1 1,2
// 2,0 2,1 2,2  2,0 2,1 2,2  2,0 2,1 2,2  2,0 2,1 2,2
//              0,0 0,1 0,2
//              1,0 1,1 1,2
//              2,0 2,1 2,2
// Each numbers pair is slice element index pair.
type Cube struct {
	n     int
	Sides [6]side
}

// String implements Stringer interface.
func (c Cube) String() string {
	n := c.n
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

func writeSideLine(str *strings.Builder, s *side, i int) {
	n := len(*s)
	for j := 0; j < n; j++ {
		str.WriteString((*s)[i][j].String() + " ")
	}
}

// NewCube generates a new cube in the solved state.
func NewCube(n int) (Cube, error) {
	if n < 1 {
		return Cube{}, errors.New("invalid cube size should be > 0")
	}

	c := Cube{n: n}
	for i := range c.Sides {
		c.Sides[i] = make([][]face, n, n)
		for j := range c.Sides[i] {
			c.Sides[i][j] = make([]face, n, n)
			for k := range c.Sides[i][j] {
				c.Sides[i][j][k] = face(i)
			}
		}
	}
	return c, nil
}

// Exec performs a move or rotation on a cube.
func (c Cube) exec(m Move) {
	f := m.face
	s := c.Sides

	s[f].turn(m.dir)
	if m.rotationFor(c) {
		s[f.opposite()].turn(m.dir.opposite())
	}

	switch f {
	case U:
		for i := 0; i < m.N; i++ {
			for j := 0; j < c.n; j++ {
				cycle(m.dir, &s[F][i][j], &s[L][i][j], &s[B][i][j], &s[R][i][j])
			}
		}
	case D:
		for i := 0; i < m.N; i++ {
			for j := 0; j < c.n; j++ {
				cycle(m.dir, &s[F][c.n-1-i][j], &s[U][c.n-1-i][j], &s[B][c.n-1-i][j], &s[D][c.n-1-i][j])
			}
		}
	case R:
		for i := 0; i < m.N; i++ {
			for j := 0; j < c.n; j++ {
				cycle(m.dir, &s[F][j][c.n-1-i], &s[U][j][c.n-1-i], &s[B][j][i], &s[D][j][c.n-1-i])
			}
		}
	case L:
		for i := 0; i < m.N; i++ {
			for j := 0; j < c.n; j++ {
				cycle(m.dir, &s[F][j][i], &s[U][j][i], &s[B][c.n-1-j][c.n-1-i], &s[D][j][i])
			}
		}
	case F:
		for i := 0; i < m.N; i++ {
			for j := 0; j < c.n; j++ {
				cycle(m.dir, &s[U][c.n-1-i][j], &s[R][j][i], &s[D][i][c.n-1-j], &s[L][c.n-1-j][c.n-1-i])
			}
		}
	case B:
		for i := 0; i < m.N; i++ {
			for j := 0; j < c.n; j++ {
				cycle(m.dir, &s[U][i][j], &s[R][j][c.n-1-i], &s[D][c.n-1-i][c.n-1-j], &s[L][c.n-1-j][i])
			}
		}
	}
}
