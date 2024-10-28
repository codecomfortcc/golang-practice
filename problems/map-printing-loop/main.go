package main

import "fmt"

func main() {
	// m:= map[string]string{}
	m:= make(map[string]string)
	m["name"]= "john"
	m["age"]= "20"
	m["city"]= "New York"
	for k,v := range m{
		fmt.Println(k,":",v)
	}
}
