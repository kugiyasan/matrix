package main

import (
	"fmt"
	"time"
)

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

func timeitRepeat(stmt func(), number, repeat int) []time.Duration {
	times := make([]time.Duration, repeat)

	for r := 0; r < repeat; r++ {
		start := time.Now()
		for n := 0; n < number; n++ {
			stmt()
		}
		times[r] = time.Since(start)
	}
	return times
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
	// print(A.Det())
	// print(A.Tr())

	// print(H.Minor(2, 2))
	// print(H.Cofactor(2, 2))
	// print(H.Det())

	number := 100000
	repeat := 5
	fmt.Println(timeitRepeat(matrixBenchmark, number, repeat))
	// fmt.Println(timeitRepeat(forRange, 1000, repeat))
	// fmt.Println(timeitRepeat(forClassic, 1000, repeat))
}
