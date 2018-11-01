package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func ret() (int, string) {
	return 4, "nessaa"
}

func main() {
	nums := []int{1, 2, 3, 4}
	for _, num := range nums {
		fmt.Println(num)
	}
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	fmt.Println(add(2, 4))
	fmt.Println(sub(2, 4))
	a, b := ret()
	fmt.Println(a, b)
}
