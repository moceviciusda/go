package triangle

type Kind int

const (
	NaT = -1 // not a triangle
	Equ = 0  // equilateral
	Iso = 1  // isosceles
	Sca = 2  // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b < c || b+c < a || c+a < b {
		return NaT
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || a == c || b == c {
		return Iso
	}

	return Sca
}
