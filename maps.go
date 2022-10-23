package main

import "fmt"

func main() {
	// numbers := make(map[int]string)
	// elements[1] = "Satu"
	// elements[2] = "Dua"
	// elements[3] = "Tiga"
	// elements[4] = "Empat"
	// elements[5] = "Lima"

	// elements := map[int]string{
	// 	1: "Satu",
	// 	2: "Dua",
	// 	3: "Tiga",
	// 	4: "Empat",
	// 	5: "Lima",
	// }

	numbers := map[int]map[string]string{
		1: map[string]string{
			"alias": "Satu",
		},
		2: map[string]string{
			"alias": "Dua",
		},
		3: map[string]string{
			"alias": "Tiga",
		},
		4: map[string]string{
			"alias": "Empat",
		},
		5: map[string]string{
			"alias": "Lima",
		},
	}

	for i, number := range numbers {
		if angka, ok := number["alias"]; ok {
			fmt.Println(i, angka, ok)
		}
	}

	fmt.Println(numbers)
}
