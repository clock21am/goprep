package main

import "os"

func main() {
	panic("it is problem")

	_, err := os.Create("/tmp/ttt")
	if err != nil {
		panic(err)
	}
	panic("not a prblem")
}
