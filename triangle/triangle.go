package triangle

import "math"

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

func KindFromSides(a, b, c float64) Kind {

	switch {

	case a <= 0.0 || b <= 0.0 || c <= 0.0:
		return NaT
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return NaT
	case math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0):
		return NaT
	case a+b < c || b+c < a || c+a < b:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || b == c || a == c:
		return Iso
	case a != b && b != c && a != c:
		return Sca
	}
	return Sca
}
