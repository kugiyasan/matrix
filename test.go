package main

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
	A := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	B := Matrix{
		{1, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
	}
	E := Matrix{
		{1, 2},
		{2, -3},
	}

	MatrixAdd(A, B)
	MatrixSub(A, B)
	MatrixMul(6, A)
	DotProduct(A, B)
	DotProduct(B, A)
	A.Minor(2, 2)
	A.Cofactor(2, 2)
	E.Inv()
	A.T()
	A.Det()
	A.Tr()
}
