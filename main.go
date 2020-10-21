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
		case error:
			if t != nil {
				fmt.Println(t)
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
	repeat := 5
	fmt.Println(timeitRepeat(matrixBenchmark, 10000, repeat))
}
