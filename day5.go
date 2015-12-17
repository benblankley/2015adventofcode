package main

import (
	"bufio"
	"fmt"
	"os"
//	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day5input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line is: ", line)
		fmt.Println(strings.ContainsAny(line, "a & e & i & o & u"))
		if (strings.Contains(line, "a") && strings.Contains(line, "e") && strings.Contains(line, "i")) {
			fmt.Println("Contains aei")
		} 
		if (strings.Contains(line, "a") && strings.Contains(line, "e") && strings.Contains(line, "o")) {
			fmt.Println("Contains aeo")
		} 
		if (strings.Contains(line, "a") && strings.Contains(line, "e") && strings.Contains(line, "u")) {
			fmt.Println("Contains aeu")
		} 
		if (strings.Contains(line, "e") && strings.Contains(line, "i") && strings.Contains(line, "o")) {
			fmt.Println("Contains eio")
		}
		if (strings.Contains(line, "e") && strings.Contains(line, "i") && strings.Contains(line, "u")) {
			fmt.Println("Contains eiu")
		}
		if (strings.Contains(line, "i") && strings.Contains(line, "o") && strings.Contains(line, "u")) {
			fmt.Println("Contains iou")
		}

	}
}
