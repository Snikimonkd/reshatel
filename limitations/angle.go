package limitations

import "math"

type AngleLimitation struct {
	amountOfEquations int
	amountOfDots      int
	dots              []int
	angle             float64
}

func NewAngleLimitation(dots []int, angle float64) AngleLimitation {
	return AngleLimitation{
		amountOfEquations: 1,
		amountOfDots:      4,
		dots:              dots,
		angle:             angle,
	}
}

func (f AngleLimitation) AmountOfEquations() int {
	return f.amountOfEquations
}

func (f AngleLimitation) AmountOfDots() int {
	return f.amountOfDots
}

func (f AngleLimitation) PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64) {
	// проставить лямды в нижнем правом уголочке
	m := len(lyambdas)

	cord1 := f.dots[0]
	cord2 := f.dots[1]
	cord3 := f.dots[2]
	cord4 := f.dots[3]

	cosa := math.Cos(f.angle * math.Pi / 180)

	a := dots[cord2].X - dots[cord1].X
	b := dots[cord2].Y - dots[cord1].Y
	c := dots[cord4].X - dots[cord3].X
	d := dots[cord4].Y - dots[cord3].Y

	A := (a*c*c + b*c*c) - a*(c*c+d*d)*cosa*cosa
	B := (b*d*d + a*c*d) - b*(c*c+d*d)*cosa*cosa
	C := (c*a*a + a*b*d) - c*(a*a+b*b)*cosa*cosa
	D := (d*b*b + a*b*c) - d*(a*a+b*b)*cosa*cosa

	// края
	matrix[m+cord1*2][equationsCounter] += 2 * D
	matrix[m+cord1*2+1][equationsCounter] += 2 * C

	matrix[m+cord2*2][equationsCounter] += -2 * D
	matrix[m+cord2*2+1][equationsCounter] += -2 * C

	matrix[m+cord3*2][equationsCounter] += 2 * B
	matrix[m+cord3*2+1][equationsCounter] += 2 * A

	matrix[m+cord4*2][equationsCounter] += -2 * B
	matrix[m+cord4*2+1][equationsCounter] += -2 * A

	// нижний правый угол

	// ставим значения в правую часть
	rightPart[m+cord1*2] += -2 * D * lyambdas[equationsCounter]
	rightPart[m+cord1*2+1] += -2 * C * lyambdas[equationsCounter]

	rightPart[m+cord2*2] += 2 * D * lyambdas[equationsCounter]
	rightPart[m+cord2*2+1] += 2 * C * lyambdas[equationsCounter]

	rightPart[m+cord3*2] += -2 * B * lyambdas[equationsCounter]
	rightPart[m+cord3*2+1] += -2 * A * lyambdas[equationsCounter]

	rightPart[m+cord4*2] += 2 * B * lyambdas[equationsCounter]
	rightPart[m+cord4*2+1] += 2 * A * lyambdas[equationsCounter]

	rightPart[equationsCounter] += -(a*a*c*c + 2*a*b*c*d + b*b*d*d - (a*a+b*b)*(c*c+d*d)*cosa*cosa)

	return matrix, rightPart
}
