package limitations

type FixLimitation struct {
	amountOfEquations int
	amountOfDots      int
	dots              []int
}

func NewFixLimitation(dots []int) FixLimitation {
	return FixLimitation{
		amountOfEquations: 2,
		amountOfDots:      1,
		dots:              dots,
	}
}

func (f FixLimitation) AmountOfEquations() int {
	return f.amountOfEquations
}

func (f FixLimitation) AmountOfDots() int {
	return f.amountOfDots
}

func (f FixLimitation) PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64) {
	// проставить лямды в нижнем правом уголочке
	m := len(lyambdas)

	cord1 := f.dots[0]

	// края
	matrix[m+cord1*2][equationsCounter] += 1
	matrix[m+cord1*2+1][equationsCounter+1] += 1

	// ставим значения в правую часть
	// rightPart[equationsCounter] += dots[cord1].X
	// rightPart[equationsCounter+1] += dots[cord1].Y

	return matrix, rightPart
}
