package main

import "fmt"

func main() {
	prices := [4]float64 {10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(prices[2])

	var productNames [4]string = [4]string{"A Book"}
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	featuredPrices := prices[1:3]
	featuredPrices[0] = 199.99
	fmt.Println(featuredPrices)

	fmt.Println(prices[:3])
	fmt.Println(prices[1:])

	fmt.Println(len(featuredPrices), cap(featuredPrices))

	dynPrices := []float64{10.99, 8.99}
	fmt.Println(dynPrices[0:1])

	updatedDynPrices:= append(dynPrices, 10.99, 5.99)
	fmt.Println(dynPrices)
	fmt.Println(updatedDynPrices)

	newPrices := []float64 {10.99, 9.99, 5.99, 12.99, 29.99, 100.10}
	discountPrices := []float64 {101.99, 80.99, 20.59}
	newPrices = append(newPrices, discountPrices...)
	fmt.Println(newPrices)
}
