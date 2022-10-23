package main

import "fmt"

type People struct {
	Name, Address string
	Age, Phone    int
}

func (people *People) FullName() {
	people.Name = "Riski " + people.Name
}

func main() {
	people := People{"Makira", "Aceh, Indonesia", 22, 85211234567}
	people.FullName()

	fmt.Println(people)
}
