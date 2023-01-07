package limitations

import "fmt"

// LimitationInterface - интерфейс ограничения
type LimitationInterface interface {
	// AmountOfEquations - количество уравнений
	AmountOfEquations() int
	// AmountOfDots - количество точек, в уравнениях
	AmountOfDots() int
	// PlaceValues - вставляет свои значения в матрицу
	PlaceValues(matrix [][]float64, rightPart []float64, dots []Dot, lyambdas []float64, amountOfDots int, equationsCounter int) ([][]float64, []float64)
}

// точка
type Dot struct {
	X, Y float64
}

func Printmatrix(matrix [][]float64) {
	for i := range matrix {
		for j := range matrix {
			fmt.Printf("%.2f|", matrix[i][j])
		}
		fmt.Print("\n")
	}
}
