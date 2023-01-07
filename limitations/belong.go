package limitations

type BelongLimitation struct {
	amountOfEquations int
	amountOfDots      int
	dots              []int
}

func NewBelongLimitation(dots []int) BelongLimitation {
	return BelongLimitation{
		amountOfEquations: 1,
		amountOfDots:      3,
		dots:              dots,
	}
}

func (f BelongLimitation) AmountOfEquations() int {
	return f.amountOfEquations
}

func (f BelongLimitation) AmountOfDots() int {
	return f.amountOfDots
}

func (f BelongLimitation) PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64) {
	// проставить лямды в нижнем правом уголочке
	m := len(lyambdas)

	cord1 := f.dots[0]
	cord2 := f.dots[1]
	cord3 := f.dots[2]

	// края
	matrix[m+cord1*2][equationsCounter] += (dots[cord3].Y - dots[cord1].Y) + (dots[cord1].Y - dots[cord2].Y)
	matrix[m+cord1*2+1][equationsCounter] += -(dots[cord1].X - dots[cord2].X) - (dots[cord3].X - dots[cord1].X)

	matrix[m+cord2*2][equationsCounter] += -(dots[cord3].Y - dots[cord1].Y)
	matrix[m+cord2*2+1][equationsCounter] += (dots[cord3].X - dots[cord1].X)

	matrix[m+cord3*2][equationsCounter] += -(dots[cord1].Y - dots[cord2].Y)
	matrix[m+cord3*2+1][equationsCounter] += (dots[cord1].X - dots[cord2].X)

	// нижний правый угол
	matrix[m+cord2*2+1][m+cord1*2] += -lyambdas[equationsCounter]
	matrix[m+cord3*2+1][m+cord1*2] += lyambdas[equationsCounter]

	matrix[m+cord2*2][m+cord1*2+1] += lyambdas[equationsCounter]
	matrix[m+cord3*2][m+cord1*2+1] += -lyambdas[equationsCounter]

	matrix[m+cord3*2+1][m+cord2*2] += -lyambdas[equationsCounter]
	matrix[m+cord3*2][m+cord2*2+1] += lyambdas[equationsCounter]

	// ставим значения в правую часть
	rightPart[equationsCounter] += -((dots[cord1].X-dots[cord2].X)*(dots[cord3].Y-dots[cord1].Y) - (dots[cord1].Y-dots[cord2].Y)*(dots[cord3].X-dots[cord1].X))

	return matrix, rightPart
}
