package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	s := person{name: "nessaa", age: 40}
	fmt.Println(s)
}
