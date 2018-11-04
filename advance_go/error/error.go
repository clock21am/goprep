package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("bla bla bla bla")
	}
	return arg + 2, nil
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: sqare negative number")
	}
	// implementations
	return 1, nil
}
func main() {
	for _, i := range []int{7, 42} {
		fmt.Println(i)
		if r, e := f1(i); e != nil {
			fmt.Println("error")
		} else {
			fmt.Println("no error happened", r)
		}
	}
	f, err := os.Open("Bobo.txt")
	if err != nil {
		log.Fatal(err, f)
	}
	t, err := sqrt(-1)
	if err != nil {
		fmt.Println(err, t)
	}

}
