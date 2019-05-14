package cube

// cycle performs a 4-cycle of given stickers
func cycle(dir dir, a, b, c, d *face) {
	switch dir {
	case Cw:
		cycleCw(a,b,c,d)
	case Acw:
		cycleAcw(a,b,c,d)
	case Two:
		swap(a, c)
		swap(b, d)
	}
}

func cycleCw(a, b, c, d *face) {
	*a, *b, *c, *d = *d, *a, *b, *c
}

func cycleAcw(a, b, c, d *face) {
	*a, *b, *c, *d = *b, *c, *d, *a
}

func swap(a, b *face) {
	*a, *b = *b, *a
}
