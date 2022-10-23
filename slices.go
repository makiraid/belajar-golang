package main

import "fmt"

func main() {
	// x := make([]float64, 5)
	// x := make([]float64, 5, 10)
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:5]
	fmt.Println(x)

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 3)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
