package main

import "fmt"

func main() {
	// slice print day
	slice := []string{}
	slice = append(slice, "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")
	sliceTwo:= make([]string, 3)
	sliceTwo[0] = "Monday"
	sliceTwo[1] = "Tuesday"
	sliceTwo[2] = "Wednesday"
	sliceTwo=append(sliceTwo, "Thursday", "Friday", "Saturday", "Sunday")
	loop(slice)
	loop(sliceTwo)
}
func loop(slice []string){
	for _,v := range slice{
		fmt.Println(v)
	}
}
