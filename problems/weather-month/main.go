package main

import "fmt"

func main() {
	fmt.Println("Enter month name:")
	var month string
	fmt.Scanf("%s", &month)
	switch month {
		case "January","February" , "December": fmt.Println("Winter")
		case "March","April","May": fmt.Println("Spring")
		case "June","July","August": fmt.Println("Summer")
		case "September","October","November": fmt.Println("Autumn")
		default: fmt.Println("Invalid month")
	}
	
}
