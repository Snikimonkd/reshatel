package main

import (
	"errors"
	"fmt"
	"math"
	"reshatel/limitations"
)

const eps = 0.000001

func CreateMatrix(limitationsArr []limitations.LimitationInterface, dots []limitations.Dot) ([]limitations.Dot, error) {
	amountOfEquations := 0
	amountOfDots := len(dots)
	for _, limitation := range limitationsArr {
		amountOfEquations += limitation.AmountOfEquations()
	}

	matrixSize := amountOfDots*2 + amountOfEquations

	lyambdas := make([]float64, amountOfEquations, amountOfEquations)
	asdf := 0
	for {
		// asdf++
		// if asdf > 5 {
		// 	break
		// }
		_ = asdf
		rightPart := make([]float64, matrixSize, matrixSize)
		matrix := make([][]float64, matrixSize, matrixSize)
		for i := range matrix {
			matrix[i] = make([]float64, matrixSize, matrixSize)
		}

		// поставить единички на диагонали в нижнем правом углу
		for i := (matrixSize - amountOfDots*2); i < matrixSize; i++ {
			matrix[i][i] = 1
		}

		equationsCounter := 0
		// ограничения лепят свои коэффициенты
		for _, limitation := range limitationsArr {
			matrix, rightPart = limitation.PlaceValues(matrix, rightPart, dots, lyambdas, amountOfDots, equationsCounter)
			equationsCounter += limitation.AmountOfEquations()
		}

		// копируем матрицу относительно диагонали
		for i := 0; i < matrixSize; i++ {
			for j := i + 1; j < matrixSize; j++ {
				matrix[i][j] = matrix[j][i]
			}
		}

		limitations.Printmatrix(matrix)
		fmt.Println("right part: ", rightPart)

		deltas, ok := solve(matrix, rightPart)
		if !ok {
			return nil, errors.New("wtf")
		}

		fmt.Println("deltas:", deltas)

		for i := 0; i < len(lyambdas); i++ {
			lyambdas[i] += deltas[i]
		}

		for i := 0; i < len(dots); i++ {
			dots[i].X += deltas[len(lyambdas)+i*2]
			dots[i].Y += deltas[len(lyambdas)+i*2+1]
		}

		flag := true
		for i := 0; i < len(deltas); i++ {
			if math.Abs(deltas[i]) > eps {
				flag = false
			}
		}

		fmt.Println("dots:", dots)
		fmt.Println("lymbdas:", lyambdas)
		fmt.Println("------------------------------------------")

		if flag {
			break
		}
	}

	return dots, nil
}

func PrintDots(dots []limitations.Dot) {
	for i := 0; i < len(dots); i++ {
		fmt.Println(dots[i].X, dots[i].Y)
	}
}

func main() {
	dots := []limitations.Dot{{X: 1, Y: 5}, {X: 2, Y: -12}, {X: 32, Y: -9}}

	limitations := []limitations.LimitationInterface{
		// limitations.NewDistanceLimitation([]int{0, 1}, 6),
		// limitations.NewFixLimitation([]int{0}),
		// limitations.NewVerticalLimitation([]int{0, 1}),
		// limitations.NewGorizontalLimitation([]int{0, 2}),
		// limitations.NewOverlapLimitation([]int{2, 3}),
		// limitations.NewFixLimitation([]int{0}),
		// limitations.NewVerticalLimitation([]int{0, 1}),
		// limitations.NewDistanceLimitation([]int{0, 1}, 6),
		// limitations.NewGorizontalLimitation([]int{0, 1}),
		// limitations.NewPerpendicularLimitation([]int{0, 1, 2, 3}),
		// limitations.NewDistanceLimitation([]int{0, 1}, 6),
		limitations.NewBelongLimitation([]int{0, 1, 2}),
		// limitations.NewVerticalLimitation([]int{0, 1}),
	}

	dots, err := CreateMatrix(limitations, dots)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("result:", dots)
}

type pair struct {
	i int
	j int
}

func solve(A [][]float64, B []float64) ([]float64, bool) {
	swap := []pair{}
	x := []float64{}
	for i := 0; i < len(A); i++ {
		if A[i][i] == 0 {
			// Делаем перестановку столбцов, подставляем наибольший коэффициент в диоганальный элемент
			var notZeroElementIndex int = 0
			for j := i; j < len(A[i]); j++ {
				if math.Abs(A[i][j]) > 0 {
					notZeroElementIndex = j
					break // находим первый не нулевой элемент, не будем проходить все элементы
				}
			}
			if notZeroElementIndex != 0 {
				swapColumns(i, notZeroElementIndex, A)
				swap = append(swap, pair{i, notZeroElementIndex})
			}
		}

		for l := 0; l < len(A); l++ {
			if l != i { // проходимся по остальным строкам
				var k float64 = A[l][i] / A[i][i]
				for j := i; j < len(A[i]); j++ {
					A[l][j] = A[l][j] - A[i][j]*k
				}
				B[l] = B[l] - B[i]*k
			}
		}
	}

	for i := len(A) - 1; i >= 0; i-- {
		var summ float64 = 0
		for j := len(A) - 1; j > i; j-- {
			if math.IsNaN(A[i][j]) {
				A[i][j] = 0
			}
			summ = summ + A[i][j]
		}
		x = append(x, (B[i]-summ)/A[i][i])
		for l := i; l >= 0; l-- {
			A[l][i] = A[l][i] * (B[i] - summ) / A[i][i]
		}
	}

	for i := 0; i < len(x)/2; i++ {
		x[i], x[len(x)-1-i] = x[len(x)-1-i], x[i]
	}

	for i := len(swap) - 1; i >= 0; i-- {
		x[swap[i].i], x[swap[i].j] = x[swap[i].j], x[swap[i].i]
	}

	for i := 0; i < len(x); i++ {
		if math.IsNaN(x[i]) {
			x[i] = 0
		}
	}

	return x, true
}

func swapColumns(j int, notZeroElementIndex int, A [][]float64) {
	for i := 0; i < len(A); i++ {
		A[i][j], A[i][notZeroElementIndex] = A[i][notZeroElementIndex], A[i][j]
	}
}
