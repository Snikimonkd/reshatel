package limitations

type ParallelLimitation struct {
	amountOfEquations int
	amountOfDots      int
	dots              []int
}

func NewParallelLimitation(dots []int) ParallelLimitation {
	return ParallelLimitation{
		amountOfEquations: 1,
		amountOfDots:      4,
		dots:              dots,
	}
}

func (f ParallelLimitation) AmountOfEquations() int {
	return f.amountOfEquations
}

func (f ParallelLimitation) AmountOfDots() int {
	return f.amountOfDots
}

func (f ParallelLimitation) PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64) {
	// проставить лямды в нижнем правом уголочке
	m := len(lyambdas)

	cord1 := f.dots[0]
	cord2 := f.dots[1]
	cord3 := f.dots[2]
	cord4 := f.dots[3]

	// края
	matrix[m+cord1*2][equationsCounter] += -(dots[cord4].Y - dots[cord3].Y)
	matrix[m+cord1*2+1][equationsCounter] += (dots[cord4].X - dots[cord3].X)

	matrix[m+cord2*2][equationsCounter] += (dots[cord4].Y - dots[cord3].Y)
	matrix[m+cord2*2+1][equationsCounter] += -(dots[cord4].X - dots[cord3].X)

	matrix[m+cord3*2][equationsCounter] += (dots[cord2].Y - dots[cord1].Y)
	matrix[m+cord3*2+1][equationsCounter] += -(dots[cord2].X - dots[cord1].X)

	matrix[m+cord4*2][equationsCounter] += -(dots[cord2].Y - dots[cord1].Y)
	matrix[m+cord4*2+1][equationsCounter] += (dots[cord2].X - dots[cord1].X)

	// нижний правый угол
	matrix[m+cord3*2+1][m+cord1*2] += lyambdas[equationsCounter]
	matrix[m+cord4*2+1][m+cord1*2] += -lyambdas[equationsCounter]

	matrix[m+cord3*2][m+cord1*2+1] += -lyambdas[equationsCounter]
	matrix[m+cord4*2][m+cord1*2+1] += lyambdas[equationsCounter]

	matrix[m+cord3*2+1][m+cord2*2] += lyambdas[equationsCounter]
	matrix[m+cord4*2+1][m+cord2*2] += -lyambdas[equationsCounter]

	matrix[m+cord3*2][m+cord2*2+1] += -lyambdas[equationsCounter]
	matrix[m+cord4*2][m+cord2*2+1] += lyambdas[equationsCounter]

	// ставим значения в правую часть
	rightPart[m+cord1*2] += (dots[cord4].Y - dots[cord3].Y) * lyambdas[equationsCounter]
	rightPart[m+cord1*2+1] += -(dots[cord4].X - dots[cord3].X) * lyambdas[equationsCounter]

	rightPart[m+cord2*2] += -(dots[cord4].Y - dots[cord3].Y) * lyambdas[equationsCounter]
	rightPart[m+cord2*2+1] += (dots[cord4].X - dots[cord3].X) * lyambdas[equationsCounter]

	rightPart[m+cord3*2] += -(dots[cord2].Y - dots[cord1].Y) * lyambdas[equationsCounter]
	rightPart[m+cord3*2+1] += (dots[cord2].X - dots[cord1].X) * lyambdas[equationsCounter]

	rightPart[m+cord4*2] += (dots[cord2].Y - dots[cord1].Y) * lyambdas[equationsCounter]
	rightPart[m+cord4*2+1] += -(dots[cord2].X - dots[cord1].X) * lyambdas[equationsCounter]

	rightPart[equationsCounter] += -((dots[cord2].X-dots[cord1].X)*(dots[cord4].Y-dots[cord3].Y) - (dots[cord4].X-dots[cord3].X)*(dots[cord2].Y-dots[cord1].Y))

	return matrix, rightPart
}
