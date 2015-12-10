package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	file, err := os.Open("./day2input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()
	
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("line is %[1]T %[1]q\n", line)
    }
	
}