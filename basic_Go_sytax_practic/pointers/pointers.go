package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(p *int) {
	*p = 0
}

func main() {
	i := 1
	fmt.Println("intial", i)
	zeroval(i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
