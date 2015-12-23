package main

import (
//	"bufio"
	"fmt"
	"os"
	//	"regexp"
//	"strconv"
//	"strings"
)


func main() {
	file, err := os.Open("./day7input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println("--- Day 7 ---")

	// Main Loop
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%q ", line)

	}
}
