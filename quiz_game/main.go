package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	csvFile, _ := os.Open("data_set.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var count int
	count = 0
	for {
		line, error := reader.Read()
		var e int
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		fmt.Println(line[0])
		answer := line[1]
		fmt.Scanf("%d", &e)
		ans, _ := strconv.Atoi(answer)
		if ans == e {
			count++
		}
	}
	if count > 5 {
		fmt.Println("yes you have passed")
	} else {
		fmt.Println("Sorry we can't proceed with you right now")
	}
}
