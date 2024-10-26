package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	value := 3
	result := removeFirst(slice, value)
	fmt.Println("result:", result)
}
func removeFirst(slice []int, value int) []int {
	result := []int{}
	for i := 0; i < len(slice); i++ {
		if slice[i] != value {
			result = append(result, slice[i])
		}
	}
	return result
}
