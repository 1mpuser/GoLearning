package main

import (
	"fmt"
	"math"
	"sort"

	// "os"
	// geometry "/Geometry/geometry.go"
	"reflect"
)

type Point struct {
	x float64
	y float64
}

type Rectangle struct {
	a Point
	b Point
	c Point
	d Point
}

type Triangle struct {
	a Point
	b Point
	c Point
}
type Animal struct {
	voice string
}

type Robot struct {
	formFactor string
}

type RoboDog struct {
	Animal
	Robot
}

func (animal Animal) Speak() {
	fmt.Println(animal.voice)
}

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

//problem from codewars
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
			continue
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
func differenceBetweenPoints(a Point, b Point) float64 {
	var difX = math.Abs(a.x - b.x)
	var difY = math.Abs(a.y - b.y)
	return math.Sqrt(math.Pow(difX, 2) + math.Pow(difY, 2))
}

func (triangle Triangle) area() float64 {
	// tmpVar := 1 / 2
	return (triangle.a.x*(triangle.b.y-triangle.c.y) + triangle.b.x*(triangle.c.y-triangle.a.y) + triangle.c.x*(triangle.a.y-triangle.b.y)) / 2
}

func (rect Rectangle) area() float64 {
	var LengthValueArray [3]float64
	LengthValueArray[0] = differenceBetweenPoints(rect.a, rect.b)
	LengthValueArray[1] = differenceBetweenPoints(rect.a, rect.c)
	LengthValueArray[2] = differenceBetweenPoints(rect.a, rect.d)
	sort.Float64s(LengthValueArray[0:3])
	return LengthValueArray[0] * LengthValueArray[1]
}

func (triangle Triangle) perimeter() float64 {
	return differenceBetweenPoints(triangle.a, triangle.b) + differenceBetweenPoints(triangle.a, triangle.c) + differenceBetweenPoints(triangle.b, triangle.c)
}

func (rect Rectangle) perimeter() float64 {
	var LengthValueArray [3]float64
	LengthValueArray[0] = differenceBetweenPoints(rect.a, rect.b)
	LengthValueArray[1] = differenceBetweenPoints(rect.a, rect.c)
	LengthValueArray[2] = differenceBetweenPoints(rect.a, rect.d)
	sort.Float64s(LengthValueArray[0:3])
	return 2 * (LengthValueArray[0] + LengthValueArray[1])
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(reflect.TypeOf(i))
		fmt.Println(n, ":", i)
	}
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
	//fmt.Println(quote.Go())
	/////structs and interfaces
	//_________Interfaces_______________

	a := Point{0, 1}
	b := Point{2, 3}
	abc := Triangle{Point{0, 2}, Point{3, 4}, Point{5, 6}}
	fmt.Println(differenceBetweenPoints(a, b))
	fmt.Println("Triangle : ", abc.a, ",", abc.b, ",", abc.c)
	fmt.Println("Triangle square :", abc.area())
	i := 0
	fmt.Println(reflect.TypeOf(i))
	go f(0)
	var input string
	fmt.Scanln(&input)

}
