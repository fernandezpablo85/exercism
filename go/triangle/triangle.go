package triangle

type Kind int

const (
	Equ Kind = iota
	Iso
	Sca
	NaT
)

func KindFromSides(a float64, b float64, c float64) Kind {
	if !valid(a, b, c) {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}

func valid(a float64, b float64, c float64) bool {
	return (a+b) > c && (a+c) > b && (c+b) > a
}
