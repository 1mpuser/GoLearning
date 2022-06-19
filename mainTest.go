package main

import (
	"fmt"
	"math"
	// "os"
)

func averageWithSquares(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total += math.Pow(v, 2)
	}
	return total / float64(len(arr))
}

func half(a int) (int, bool) {
	if a%2 == 0 {
		return a / 2, true
	} else {
		return a / 2, false
	}
}

func findMinInArr(arr []int) int {
	min := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}

// func makeEvenGenerator() func () even{
// 	i:= even(0)
// }

func findMax(args ...int) (int, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("There is no args in function")
	}
	counter := 0
	firstValue := 0
	for _, v := range args {
		if counter == 0 {
			firstValue = v
		}
		if counter != 0 {
			break
		}
		counter++
	}
	result := firstValue
	for _, i := range args {
		if result < i {
			result = i
		}
	}
	return result, nil
}

func TwoOldestAges(ages []int) [2]int {
	var firstAge = ages[0]
	var secondAge = ages[1]
	if firstAge < secondAge {
		secondAge, firstAge = firstAge, secondAge
	}
	for i := 2; i < len(ages); i++ {
		var age = ages[i]
		fmt.Println("Younger Age : ", secondAge, " Elder Age : ", firstAge)
		fmt.Println(age)
		if age > firstAge {
			firstAge, secondAge = age, firstAge
		} else if age > secondAge {
			secondAge = age
		} else {
			continue
		}
		fmt.Println(" ")
	}
	var arr [2]int
	arr[0] = secondAge
	arr[1] = firstAge
	return arr
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
	// arr := []float64{98, 93, 77, 82, 83}
	// fmt.Println(averageWithSquares(arr))
	//замыкание = closure
	// x := make([]int, 3, 9)
	// fmt.Println((x))
	// fmt.Println(os.Args)
	// x := [6]string{"a", "b", "c", "d", "e", "f"}
	// fmt.Println(x[2:5])
	//last number excluded
	// x := []int{
	// 	48, 96, 86, 68,
	// 	57, 82, 63, 70,
	// 	37, 34, 83, 27,
	// 	19, 97, 9, 17,
	// }
	// fmt.Println(findMinInArr(x))
	// for i := 0; i < len(x); i++ {
	// 	str := ' '
	// 	fmt.Print(x[i])
	// 	fmt.Print(str)
	// 	fmt.Println(half(x[i]))
	// }
	x := [10]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}
	fmt.Println(TwoOldestAges(x[0:len(x)]))
}
