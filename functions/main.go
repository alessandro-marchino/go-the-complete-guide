package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	moreNumbers := []int{5, 1, 2}
	transformedNumbers1 := transformNumbers(&numbers, getTransformerFunction(&numbers))
	transformedNumbers2 := transformNumbers(&moreNumbers, getTransformerFunction(&moreNumbers))
	fmt.Println(transformedNumbers1)
	fmt.Println(transformedNumbers2)

	transformed := transformNumbers(&numbers, func(num int) int { return num * 2 })
	fmt.Println(transformed)

	transformed2 := transformNumbers(&numbers, createTransformer(5))
	fmt.Println(transformed2)

	fact := factorial(5)
	fmt.Println("Factorial of 5 is", fact)

	fmt.Println("Sumup", sumup(1, 2, 3, 4))
	fmt.Println("Sumup", sumup(1, numbers...))
}

func transformNumbers(numbers *[]int, transform transformFn) (dNumbers []int) {
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	}
	return triple
}

func createTransformer(factor int) func(int) int {
	return func(num int) int {
		return num * factor
	}
}

func factorial(number int) int {
	if number <= 1 {
		return 1
	}
	return number * factorial(number - 1)
}

func sumup(startingValue int, numbers ...int) int {
	sum := startingValue
	for _, val := range numbers {
		sum += val
	}
	return sum
}
