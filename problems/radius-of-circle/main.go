package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const PI = 3.14159
	fmt.Println("Enter the radius of the circle: ")
	reader := bufio.NewReader(os.Stdin)
	radius,_,err:= reader.ReadLine()
	if err != nil {
		fmt.Println(err)
		return	
	}
	numRadius,_ := strconv.ParseFloat(strings.TrimSpace(string(radius)), 64)
	area:= PI * numRadius * numRadius
	fmt.Printf("The area of the circle is %v \n", area)
}
