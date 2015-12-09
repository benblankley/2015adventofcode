package main

import (
	"fmt"
	"os"
//	"strings"
)

func main() {
	file, err := os.Open("./day2input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	
	for {

		data := make([]byte, 9)
		_, err := file.Read(data)
		if err != nil {
			break
		}
		fmt.Printf("Package: %q\n", data)

	}
}
