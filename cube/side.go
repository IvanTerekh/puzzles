package cube

// side represents a square of stickers on each side of the cube.
type side [][]face

func (s side) turn(d dir) {
	n := len(s)
	m := n/2 + n%2
	for i := 0; i < m; i++ {
		for j := 0; j < n/2; j++ {
			cycle(d, &s[i][j], &s[j][n-1-i], &s[n-1-i][n-1-j], &s[n-1-j][i])
		}
	}
}

func (s side) isSolved() bool {
	first := s[0][0]
	for _, line := range s {
		for _, curr := range line {
			if curr != first {
				return false
			}
		}
	}
	return true
}
