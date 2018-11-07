package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func fs(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, "::", i)
	}
}

func main() {
	f("direct")

	go fs("goroutine")

	go f("thread")

	// go fmt.Scanln()

	fmt.Scanln()
}
