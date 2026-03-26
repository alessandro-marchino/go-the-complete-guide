package main

import "fmt"

func main() {
	age := 32
	agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("AgePointer:", agePointer)
	fmt.Println("Age:", *agePointer)
	fmt.Println(getAdultYears(agePointer))
}

func getAdultYears(age *int) int {
	return *age - 18
}
