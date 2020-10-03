package main

import (
	"math/rand"
)

func forRange() {
	output := [1000][1000]float64{}

	for i := range output {
		for j := range output[i] {
			output[i][j] = float64(i + j)
		}
	}
}

func forClassic() {
	output := [1000][1000]float64{}

	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]); j++ {
			output[i][j] = float64(i + j)
		}
	}
}

func matrixBenchmark() {
	// 	A := Matrix{
	// 		{1, 2, 3},
	// 		{4, 5, 6},
	// 		{7, 8, 9},
	// 	}
	// B := Matrix{
	// 	{1, 0, 0},
	// 	{0, 0, 1},
	// 	{0, 1, 0},
	// }
	// E := Matrix{
	// 	{1, 2},
	// 	{2, -3},
	// }

	A := Matrix(create2DArray(10, 10))
	for i := range A {
		for j := range A[i] {
			A[i][j] = rand.Float64()
		}
	}

	// create2DArray(10, 10)
	// I(100)
	// MatrixAdd(A, B)
	// MatrixSub(A, B)
	// DotProduct(A, B)
	// DotProduct(B, A)
	// MatrixMul(6, A)
	// A.T()
	// A.Tr()
	// A.Minor(2, 2)
	// A.Cofactor(2, 2)
	A.Det()
	A.Adj()
	// E.Inv()
}
