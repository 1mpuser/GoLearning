package main

import (
	"fmt"
	"math"
)

func averageWithSquares(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += math.Pow(v, 2)
	}
	return total / float64(len(arr))
}

func main() {
	fmt.Println("Da mihi esse optimum")
	// for i := 1; i < 100; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	// slice1 := []int{1, 2, 3}
	// slice2 := append(slice1, 4, 5)
	// fmt.Println(slice1, slice2)
	// fmt.Println(len(slice2))
	// tmpMap := make(map[string]int)
	// tmpMap["khuilo"] = 1

	// fmt.Println(tmpMap)
	// x := [6]string{"a", "b", "c", "d", "e", "f"}
	// fmt.Println(x[2:5])
	arr := []float64{98, 93, 77, 82, 83}
	fmt.Println(averageWithSquares(arr))
}
