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
	}
}
