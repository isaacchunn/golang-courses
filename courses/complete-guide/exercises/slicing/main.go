package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	//Task 1
	fmt.Println("Task 1")
	hobbies := [3]string{"Reading", "Playing", "Sleeping"}
	fmt.Println(hobbies)

	//Task 2
	//Print first element
	fmt.Println("Task 2")
	fmt.Println("First Hobby:", hobbies[0])
	hobbiesList := hobbies[1:]
	fmt.Println(hobbiesList)

	//Task 3
	fmt.Println("Task 3")
	firstSlice := hobbies[1:3]
	secondSlice := hobbies[2:]
	fmt.Println(firstSlice)
	fmt.Println(secondSlice)

	//Task 4
	fmt.Println("Task 4")
	firstSlice = hobbies[2:]
	fmt.Println(firstSlice)

	//Task 5
	fmt.Println("Task 5")
	courseGoals := []string{"Learn Go", "Go..."}
	fmt.Println(courseGoals)

	//Task 6
	fmt.Println("Task 6")
	courseGoals[1] = "Learn More Go"
	courseGoals = append(courseGoals, "MORE GO")
	fmt.Println(courseGoals)

	//Task 7
	fmt.Println("Task 7")
	products := []Product{
		newProduct("Dryer", "123", 125.0),
		newProduct("Comb", "111", 12.0),
	}
	fmt.Println(products)
	products = append(products, newProduct("omg", "555", 123.0))
	fmt.Println(products)
}

//Function to generate a product
func newProduct(title, id string, price float64) Product {
	return Product{
		title: title,
		id:    id,
		price: price,
	}
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
