package main

import "fmt"

// refactor this into Clean Code after course is complete

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
	exercise7()
}

// Classic hello world
func exercise1() {
	fmt.Println("Hello World")
}

// Say hello to my little friend
func exercise2() {
	name := "Dylan"
	fmt.Println("Hello", name)
}

// Print user's name
func exercise3() {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}

// Do modulus division on 2 user-provided numbers
func exercise4() {
	fmt.Println("Enter a small integer:")
	var small int
	fmt.Scan(&small)
	fmt.Println("Enter a larger integer:")
	var large int
	fmt.Scan(&large)
	fmt.Println(large, "mod", small, "is", large%small)
}

// Print all even integers between 1 and 100
func exercise5() {
	for i := 0; i <= 100; i += 2 {
		fmt.Println(i)
	}
}

// Classic FizzBuzz problem
func exercise6() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			}
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

// Sum of multiples of 3 or 5 below 1000
func exercise7() {
	sum := 0
	for i := 3; i <= 100; i++ {
		if i%3 == 0 {
			sum += i
			continue
		} else if i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("The sum is %v.", sum)
}
