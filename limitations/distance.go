package limitations

type DistanceLimitation struct {
	amountOfEquations int
	amountOfDots      int
	dots              []int
	distance          float64
}

func NewDistanceLimitation(dots []int, distance float64) DistanceLimitation {
	return DistanceLimitation{
		amountOfEquations: 1,
		amountOfDots:      2,
		dots:              dots,
		distance:          distance,
	}
}

func (d DistanceLimitation) AmountOfEquations() int {
	return d.amountOfEquations
}

func (d DistanceLimitation) AmountOfDots() int {
	return d.amountOfDots
}

func (d DistanceLimitation) PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64) {
	// проставить лямды в нижнем правом уголочке
	m := len(lyambdas)

	cord1 := d.dots[0]
	cord2 := d.dots[1]

	// края
	matrix[m+cord1*2][equationsCounter] += d.dF_dl_dx1(dots[cord1].X, dots[cord2].X)
	matrix[m+cord1*2+1][equationsCounter] += d.dF_dl_dy1(dots[cord1].Y, dots[cord2].Y)
	matrix[m+cord2*2][equationsCounter] += d.dF_dl_dx2(dots[cord1].X, dots[cord2].X)
	matrix[m+cord2*2+1][equationsCounter] += d.dF_dl_dy2(dots[cord1].Y, dots[cord2].Y)

	// середина
	// диагональ
	matrix[m+cord1*2][m+cord1*2] += d.lymbda(lyambdas[equationsCounter])
	matrix[m+cord1*2+1][m+cord1*2+1] += d.lymbda(lyambdas[equationsCounter])
	matrix[m+cord2*2][m+cord2*2] += d.lymbda(lyambdas[equationsCounter])
	matrix[m+cord2*2+1][m+cord2*2+1] += d.lymbda(lyambdas[equationsCounter])

	// всратые с минусами
	matrix[m+cord1*2][m+cord2*2] += -d.lymbda(lyambdas[equationsCounter])
	matrix[m+cord2*2][m+cord1*2] += -d.lymbda(lyambdas[equationsCounter])
	matrix[m+cord1*2+1][m+cord2*2+1] += -d.lymbda(lyambdas[equationsCounter])
	matrix[m+cord2*2+1][m+cord1*2+1] += -d.lymbda(lyambdas[equationsCounter])

	// ставим значения в правую часть
	rightPart[equationsCounter] += -d.dF_dl(dots[cord1].X, dots[cord1].Y, dots[cord2].X, dots[cord2].Y, d.distance)

	return matrix, rightPart
}

// правая часть

func (d DistanceLimitation) dF_dl(x1, y1, x2, y2, dist float64) float64 {
	return (x2-x1)*(x2-x1) + (y2-y1)*(y2-y1) - dist*dist
}

// лево и верх

func (d DistanceLimitation) dF_dl_dx1(x1, x2 float64) float64 {
	return -2 * (x2 - x1)
}

func (d DistanceLimitation) dF_dl_dx2(x1, x2 float64) float64 {
	return 2 * (x2 - x1)
}

func (d DistanceLimitation) dF_dl_dy1(y1, y2 float64) float64 {
	return -2 * (y2 - y1)
}

func (d DistanceLimitation) dF_dl_dy2(y1, y2 float64) float64 {
	return 2 * (y2 - y1)
}

// серединка

func (d DistanceLimitation) lymbda(lymabda float64) float64 {
	return 2 * lymabda
}
