package main

import (
	"fmt"
	"math"
	"time"
)

// Matrix has some interesting methods for matrices computation
type Matrix [][]float64

func create2DArray(x, y int) [][]float64 {
	output := make([][]float64, x)
	for i := range output {
		output[i] = make([]float64, y)
	}
	return output

	// https://stackoverflow.com/questions/23869717/initialize-a-2d-dynamic-array-in-go
	// M := make([][]uint8, row)
	// e := make([]uint8, row * col)
	// for i := range M {
	// 	a[i] = e[i * col:(i + 1) * col]
	// }
}

// I creates a identity Matrix of size n
func I(n int) Matrix {
	output := create2DArray(n, n)
	for i := range output {
		output[i][i] = 1
	}
	return output
}

// MatrixAdd adds two matrices together
func MatrixAdd(A, B Matrix) (Matrix, *string) {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		err := "Can't add two matrices of different size"
		return Matrix{}, &err
	}

	output := create2DArray(len(A), len(A[0]))

	for i := range output {
		for j := range output[0] {
			output[i][j] = A[i][j] + B[i][j]
		}
	}
	return output, nil
}

// MatrixSub substracts two matrices together
func MatrixSub(A, B Matrix) (Matrix, *string) {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		err := "Can't substract two matrices of different size"
		return Matrix{}, &err
	}

	output := create2DArray(len(A), len(A[0]))

	for i := range output {
		for j := range output[0] {
			output[i][j] = A[i][j] - B[i][j]
		}
	}
	return output, nil
}

// MatrixMul multiply the Matrix A by the scalar k
func MatrixMul(k float64, A Matrix) Matrix {
	output := create2DArray(len(A), len(A[0]))

	for i := range output {
		for j := range output[0] {
			output[i][j] = k * A[i][j]
		}
	}
	return output
}

// DotProduct does a matrix multiplication (AB)
func DotProduct(A, B Matrix) (Matrix, *string) {
	if len(A[0]) != len(B) {
		err := "Those matrices aren't compatible for dot product"
		return Matrix{}, &err
	}

	n := len(A[0])
	output := create2DArray(len(A), len(B[0]))

	for i := range output {
		for j := range output[0] {
			sum := 0.
			for k := 0; k < n; k++ {
				sum += A[i][k] * B[k][j]
			}
			output[i][j] = sum
		}
	}
	return output, nil

}

// Minor returns the minor of the Matrix
func (A Matrix) Minor(i, j int) (float64, *string) {
	if i >= len(A) || j >= len(A[0]) {
		err := fmt.Sprintf("(%d, %d) is outside the Matrix of size (%d, %d)", i, j, len(A), len(A[0]))
		return 0, &err
	}

	// Deep copy the Matrix
	output := create2DArray(len(A), len(A[0]))
	for y := range output {
		for x := range output[i] {
			output[y][x] = A[y][x]
		}
	}

	output = append(output[:i], output[i+1:]...)
	for x := range output {
		output[x] = append(output[x][:j], output[x][j+1:]...)
	}
	return Matrix(output).Det()
}

// Det returns the determinant of the square Matrix
func (A Matrix) Det() (float64, *string) {
	if len(A) != len(A[0]) {
		err := "You need a square matrix to find the determinant"
		return 0, &err
	}

	if len(A) == 1 {
		return A[0][0], nil
	} else if len(A) == 2 {
		return A[0][0]*A[1][1] - A[1][0]*A[0][1], nil
	}

	i := 0
	n := len(A)
	sum := 0.
	for k := 0; k < n; k++ {
		cofactor, err := A.Cofactor(i, k)
		if err != nil {
			return 0, err
		}
		sum += A[i][k] * cofactor
	}

	return sum, nil
}

// Cofactor returns the cofactor of the Matrix
func (A Matrix) Cofactor(i, j int) (float64, *string) {
	minor, err := A.Minor(i, j)
	return math.Pow(-1, float64(i+j)) * minor, err
}

// Inv returns the inverted Matrix
func (A Matrix) Inv() (Matrix, *string) {
	determinant, err := A.Det()
	if err != nil {
		return Matrix{}, err
	}
	if determinant == 0 {
		err := "The determinant is null, the inverse can't be compute"
		return Matrix{}, &err
	}

	adjacent, err := A.Adj()
	return MatrixMul(1/determinant, adjacent), err
}

// Adj returns the adjacent Matrix
func (A Matrix) Adj() (Matrix, *string) {
	output := create2DArray(len(A), len(A[0]))

	for i := range output {
		for j := range output[i] {
			cofactor, err := A.Cofactor(i, j)
			if err != nil {
				return Matrix{}, err
			}
			output[i][j] = cofactor
		}
	}
	return Matrix(output).T(), nil
}

// T returns the transposed Matrix
func (A Matrix) T() Matrix {
	output := create2DArray(len(A[0]), len(A))

	for i := range output {
		for j := range output[i] {
			output[i][j] = A[j][i]
		}
	}

	return output
}

// Tr returns the trace of the Matrix
func (A Matrix) Tr() (float64, *string) {
	if len(A) != len(A[0]) {
		err := "The matrix should be in a square shape"
		return 0, &err
	}

	sum := 0.
	for k := range A {
		sum += A[k][k]
	}
	return sum, nil
}

func print(i ...interface{}) {
	for _, v := range i {
		switch t := v.(type) {
		case Matrix:
			fmt.Println("Matrix [")
			for _, x := range t {
				fmt.Printf("  %v,\n", x)
			}
			fmt.Println("]")
		case *string:
			if t != nil {
				fmt.Println(*t)
			}
		default:
			fmt.Println(v)
		}
	}
}

func loop(number int) {
	for x := 0; x < number; x++ {
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
}

func repeat(repeat int, number int) {
	for x := 0; x < repeat; x++ {
		start := time.Now()

		loop(number)

		end := time.Since(start)
		fmt.Println(end)
	}
}

func main() {
	// A := Matrix{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9}}

	// B := Matrix{
	// 	{1, 0, 0},
	// 	{0, 0, 1},
	// 	{0, 1, 0}}
	// E := Matrix{
	// 	{1, 2},
	// 	{2, -3}}
	// H := Matrix{
	// 	{4, 2, 3, 9, 9},
	// 	{-2, 4, 7, -7, -7},
	// 	{2, 3, 11, 1, 1},
	// 	{1, 1, 2, -3, -1},
	// 	{1, 1, 2, 0, 1}}
	// H := Matrix{
	// 	{2, 3, 9, 9},
	// 	{4, 7, -7, -7},
	// 	{3, 11, 1, 1},
	// 	{1, 2, -3, -1}}

	// print(MatrixAdd(A, B))
	// print(MatrixSub(A, B))
	// print(MatrixMul(6, A))
	// print(DotProduct(A, B))
	// print(DotProduct(B, A))
	// print(A.Minor(2, 2))
	// print(A.Cofactor(2, 2))
	// print(A.T())
	// print(A.Det()) // wrong
	// print(A.Tr())

	// print(H.Minor(2, 2))
	// print(H.Cofactor(2, 2))
	// print(H)
	// print(H.Det())
	// print(H)

	repeat(5, 10000)
}
