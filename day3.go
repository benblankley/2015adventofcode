package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./day3input.txt")
	santax := 0
	santay := 0
	
	var Presents [100][100]int

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
		if strings.EqualFold(string(data[0]), ">") {
			santax++
		} else if strings.EqualFold(string(data[0]), "<") {
			santax--
		} else if strings.EqualFold(string(data[0]), "^") {
			santay++
		} else if strings.EqualFold(string(data[0]), "v") {
			santay--
		}
		fmt.Printf("X: %10d Y: %10d \n", santax, santay)
		Presents[santax][santay] =+ 1
	}
}
