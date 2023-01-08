package limitations

type VerticalLimitation struct {
	amountOfEquations int
	amountOfDots      int
	dots              []int
}

func NewVerticalLimitation(dots []int) VerticalLimitation {
	return VerticalLimitation{
		amountOfEquations: 1,
		amountOfDots:      2,
		dots:              dots,
	}
}

func (f VerticalLimitation) AmountOfEquations() int {
	return f.amountOfEquations
}

func (f VerticalLimitation) AmountOfDots() int {
	return f.amountOfDots
}

func (f VerticalLimitation) PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64) {
	// проставить лямды в нижнем правом уголочке
	m := len(lyambdas)

	cord1 := f.dots[0]
	cord2 := f.dots[1]

	// края
	matrix[m+cord1*2][equationsCounter] += -1
	matrix[m+cord2*2][equationsCounter] += 1

	// ставим значения в правую часть
	rightPart[m+cord1*2] += lyambdas[equationsCounter]
	rightPart[m+cord2*2] += -lyambdas[equationsCounter]
	rightPart[equationsCounter] += -(dots[cord2].X - dots[cord1].X)

	return matrix, rightPart
}
