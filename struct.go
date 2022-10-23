package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Address struct {
	Street string
	City   string
}

type Phone struct {
	Code, Number int
}

func main() {
	user := User{Name: "Riski Makira", Age: 22}
	address := Address{"Jl. Medan - Banda Aceh", "Lhokseumawe"}

	phone := Phone{}
	phone.Code = 62
	phone.Number = 85211234567

	var person = []struct {
		Name   string
		Age    int
		Code   int
		Number int
		Street string
		City   string
	}{
		{user.Name, user.Age, phone.Code, phone.Number, address.Street, address.City},
		{Name: "Nama Orang", Age: 30, Code: 62, Number: 85211234568, Street: "Jalanin Aja", City: "Kota People's Republic"},
	}

	for _, people := range person {
		fmt.Printf("Name: %s\n", people.Name)
		fmt.Printf("Age: %d\n", people.Age)
		fmt.Printf("Number: %d%d\n", people.Code, people.Number)
		fmt.Printf("Street: %s\n", people.Street)
		fmt.Printf("City: %s\n", people.City)
		fmt.Println("-----------------------------------")
	}
}
