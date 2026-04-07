package main

import "fmt"

type Product struct {
	title string
	id string
	price float64
}

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	//    Output (print) that array in the command line.
	fmt.Println("1.")
	hobbies := []string {"Reading", "Playing Games", "Sleeping"}
	fmt.Println(hobbies)
	// 2) Also output more data about that array:
	//    - The first element (standalone)
	//    - The second and third element combined as a new list
	fmt.Println("2.")
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	// 3) Create a slice based on the first element that contains
	//    the first and second elements.
	//    Create that slice in two different ways (i.e. create two slices in the end)
	fmt.Println("3.")
	slice1 := hobbies[0:2]
	slice2 := hobbies[:2]
	fmt.Println(slice1)
	fmt.Println(slice2)
	// 4) Re-slice the slice from (3) and change it to contain the second
	//    and last element of the original array.
	fmt.Println("4.")
	slice1 = slice1[1:3]
	fmt.Println(slice1)
	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	fmt.Println("5.")
	course_goals := []string {"Learning Go", "Having fun"}
	fmt.Println(course_goals)
	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	fmt.Println("6.")
	course_goals[1] = "Improving myself"
	fmt.Println(course_goals)
	course_goals = append(course_goals, "Program")
	fmt.Println(course_goals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//    dynamic list of products (at least 2 products).
	//    Then add a third product to the existing list of products.
	fmt.Println("7.")
	products := []Product {
		{ title: "Title1", id: "Id1", price: 19.99 },
		{ title: "Title2", id: "Id2", price: 29.99 },
	}
	fmt.Println(products)
	products = append(products, Product{ title: "Title3", id: "Id3", price: 39.99 })
	fmt.Println(products)
}


