package main

import (
	"errors"
	"fmt"
	"math"
)

// Matrix has some interesting methods for matrices computation
type Matrix [][]float64

func create2DArray(x, y int) [][]float64 {
	array := make([][]float64, x)
	for i := range array {
		array[i] = make([]float64, y)
	}
	return array
}

func copyMatrix(A *Matrix) [][]float64 {
	AM := create2DArray(len(*A), len((*A)[0]))
	for y := range AM {
		for x := range AM[y] {
			AM[y][x] = (*A)[y][x]
		}
	}
	return AM
}

func isUpperTriangular(A *Matrix) bool {
	for i := 0; i < len(*A)-1; i++ {
		for j := i + 1; j < len(*A); j++ {
			if (*A)[i][j] != 0 {
				return false
			}
		}
	}
	return true
}
func isLowerTriangular(A *Matrix) bool {
	for i := 1; i < len(*A); i++ {
		for j := 0; j < i+1; j++ {
			if (*A)[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

func isTriangular(A *Matrix) bool {
	return isUpperTriangular(A) || isLowerTriangular(A)
}

// I creates a identity Matrix of size n
func I(n int) Matrix {
	array := create2DArray(n, n)
	for i := range array {
		array[i][i] = 1
	}
	return array
}

// MatrixAdd adds two matrices together
func MatrixAdd(A, B Matrix) (Matrix, error) {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		err := errors.New("Can't add two matrices of different size")
		return Matrix{}, err
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
func MatrixSub(A, B Matrix) (Matrix, error) {
	if len(A) != len(B) || len(A[0]) != len(B[0]) {
		err := "Can't substract two matrices of different size"
		return Matrix{}, errors.New(err)
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
func DotProduct(A, B Matrix) (Matrix, error) {
	if len(A[0]) != len(B) {
		err := "Those matrices aren't compatible for dot product"
		return Matrix{}, errors.New(err)
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

// T returns the transposed Matrix
func (A *Matrix) T() Matrix {
	output := create2DArray(len((*A)[0]), len(*A))

	for i := range output {
		for j := range output[i] {
			output[i][j] = (*A)[j][i]
		}
	}

	return output
}

// Trace returns the trace of the Matrix
func (A *Matrix) Trace() (float64, error) {
	if len(*A) != len((*A)[0]) {
		err := "The matrix should have a square shape"
		return 0, errors.New(err)
	}

	sum := 0.
	for k := range *A {
		sum += (*A)[k][k]
	}
	return sum, nil
}

// Minor returns the minor of the Matrix
func (A *Matrix) Minor(i, j int) (float64, error) {
	if i >= len(*A) || j >= len((*A)[0]) {
		err := fmt.Sprintf("(%d, %d) is outside the Matrix of size (%d, %d)", i, j, len(*A), len((*A)[0]))
		return 0, errors.New(err)
	}

	// Deep copy the Matrix
	output := create2DArray(len(*A), len((*A)[0]))
	for y := range output {
		for x := range output[i] {
			output[y][x] = (*A)[y][x]
		}
	}

	output = append(output[:i], output[i+1:]...)
	for x := range output {
		output[x] = append(output[x][:j], output[x][j+1:]...)
	}

	tmp := Matrix(output)
	return tmp.Det(), nil
}

// Cofactor returns the cofactor of the Matrix
func (A Matrix) Cofactor(i, j int) (float64, error) {
	minor, err := A.Minor(i, j)
	return math.Pow(-1, float64(i+j)) * minor, err
}

// Det returns the determinant of a square Matrix
func (A *Matrix) Det() float64 {
	n := len(*A)
	var AM [][]float64

	if isTriangular(A) {
		AM = *A
	} else {
		AM = copyMatrix(A)

		// convert the matrix into a triangular form
		for diag := 0; diag < n; diag++ {
			for i := diag + 1; i < n; i++ {
				if AM[diag][diag] == 0 {
					AM[diag][diag] = 1e-200
				}

				scale := AM[i][diag] / AM[diag][diag]
				for j := 0; j < n; j++ {
					AM[i][j] -= (scale * AM[diag][j])
				}
			}
		}
	}

	product := 1.
	for i := 0; i < n; i++ {
		product *= AM[i][i]
	}
	return product
}

// Inv returns the inverted Matrix
func (A *Matrix) Inv() (Matrix, error) {
	determinant := A.Det()
	if determinant == 0 {
		err := "The determinant is null, the inverse can't be compute"
		return Matrix{}, errors.New(err)
	}

	adjacent, err := A.Adj()
	return MatrixMul(1/determinant, adjacent), err
}

// Adj returns the adjugated Matrix
func (A *Matrix) Adj() (Matrix, error) {
	output := create2DArray(len(*A), len((*A)[0]))

	for i := range output {
		for j := range output[i] {
			cofactor, err := A.Cofactor(i, j)
			if err != nil {
				return Matrix{}, err
			}
			output[i][j] = cofactor
		}
	}
	tmp := Matrix(output)
	return tmp.T(), nil
}
