package limitations

type OverlapLimitation struct {
	amountOfEquations int
	amountOfDots      int
	dots              []int
}

func NewOverlapLimitation(dots []int) OverlapLimitation {
	return OverlapLimitation{
		amountOfEquations: 2,
		amountOfDots:      2,
		dots:              dots,
	}
}

func (f OverlapLimitation) AmountOfEquations() int {
	return f.amountOfEquations
}

func (f OverlapLimitation) AmountOfDots() int {
	return f.amountOfDots
}

func (f OverlapLimitation) PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64) {
	// проставить лямды в нижнем правом уголочке
	m := len(lyambdas)

	cord1 := f.dots[0]
	cord2 := f.dots[1]

	// края
	matrix[m+cord1*2][equationsCounter] += -1
	matrix[m+cord2*2][equationsCounter] += 1

	matrix[m+cord1*2+1][equationsCounter+1] += -1
	matrix[m+cord2*2+1][equationsCounter+1] += 1

	// ставим значения в правую часть
	rightPart[equationsCounter] += -(dots[cord2].X - dots[cord1].X)
	rightPart[equationsCounter+1] += -(dots[cord2].Y - dots[cord1].Y)

	return matrix, rightPart
}
