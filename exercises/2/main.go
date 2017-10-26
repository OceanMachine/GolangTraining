package main

import "fmt"

func main() {
	// Exercise 1
	fmt.Println(exercise1(5))

	// Exercise 2: Implement exercise1 with a function expression
	exercise2 := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(exercise2(5))

	// Exercise 3
	fmt.Println(findGreatestInt([]int{1, 2, 5, 3, 18, 7}...))

	// Exercise 5
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

// Exercise 1: Return integer divided by two, and if it was even or odd
func exercise1(n int) (int, bool) {
	return n / 2, n%2 == 0
}

// Exercise 3: Write a function with one variadic parameter
func findGreatestInt(numbers ...int) int {
	var greatest int
	for _, n := range numbers {
		if n > greatest {
			greatest = n
		}
	}
	return greatest
}

// Exercise 4: (true && false) || (false && true) || !(false && false) = true

// Exercise 5: write foo that can be called many ways
func foo(numbers ...int) {
		fmt.Println(numbers)
}
