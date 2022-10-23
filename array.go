package main

import "fmt"

func multipleLine() int {
	var x [5]int
	x[0] = 100
	x[1] = 200
	x[2] = 300
	x[3] = 400
	x[4] = 500

	total := 0
	for _, value := range x {
		total += value
	}

	return total
}

func oneLine() float64 {
	x := [5]float64{600, 700, 800, 900, 1000}

	var total float64 = 0
	for _, value := range x {
		total += value
	}

	return total
}

func main() {
	multiple := multipleLine()
	one := oneLine()
	fmt.Println(multiple)
	fmt.Println(one)
}
