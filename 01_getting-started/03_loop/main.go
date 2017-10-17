package main

import "fmt"

func main() {
	// Print first 200 chars of unicode
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q\n", i, i, i, i)
	}
}
