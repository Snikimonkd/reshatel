package limitations

type GorizontalLimitation struct {
	amountOfEquations int
	amountOfDots      int
	dots              []int
}

func NewGorizontalLimitation(dots []int) GorizontalLimitation {
	return GorizontalLimitation{
		amountOfEquations: 1,
		amountOfDots:      2,
		dots:              dots,
	}
}

func (f GorizontalLimitation) AmountOfEquations() int {
	return f.amountOfEquations
}

func (f GorizontalLimitation) AmountOfDots() int {
	return f.amountOfDots
}

func (f GorizontalLimitation) PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64) {
	// проставить лямды в нижнем правом уголочке
	m := len(lyambdas)

	cord1 := f.dots[0]
	cord2 := f.dots[1]

	// края
	matrix[m+cord1*2+1][equationsCounter] += -1
	matrix[m+cord2*2+1][equationsCounter] += 1

	// ставим значения в правую часть
	rightPart[m+cord1*2+1] += lyambdas[equationsCounter]
	rightPart[m+cord1*2+1] += -lyambdas[equationsCounter]

	rightPart[equationsCounter] += -(dots[cord2].Y - dots[cord1].Y)

	return matrix, rightPart
}
