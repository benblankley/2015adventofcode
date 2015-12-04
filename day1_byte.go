package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./day1input.txt")
	floor := 0
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	for {
		data := make([]byte, 1)
		_, err := file.Read(data)
		if err != nil {
			break
		}
		if data[0] == '(' {
			floor++
		} else {
			floor--
		}

	}
	fmt.Printf("Floor: %d\n", floor)
}
