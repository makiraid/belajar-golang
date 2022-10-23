package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a, b := 50, 10
	fmt.Printf("before swap: a = %d b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("after swap: a = %d b = %d\n", a, b)
}
