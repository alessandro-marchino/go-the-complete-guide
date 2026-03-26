package main

import "fmt"

func main() {
	age := 32

	fmt.Println("Age:", age)
	fmt.Println("AgePointer:", &age)

	fmt.Println(getAdultYears(&age))
	fmt.Println("Age:", age)

	editAgeToAdultYears(&age)
	fmt.Println(age)
}

func getAdultYears(age *int) int {
	return *age - 18
}
func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
