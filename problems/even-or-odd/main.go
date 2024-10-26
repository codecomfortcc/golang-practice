package main

import "fmt"

func main() {
	var number int   
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%v", &number)
	if number%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
