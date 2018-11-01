package main

import "fmt"

/*import "time"
 */
func main() {
	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("sad")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
}
