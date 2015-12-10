package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
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
        fmt.Printf("line is %[1]T %[1]q ", line)
        coordinates := strings.Split(line, "x")

        l, _ := strconv.Atoi(coordinates[0])
        w, _ := strconv.Atoi(coordinates[1])
        h, _ := strconv.Atoi(coordinates[2])
        fmt.Printf("%d %d %d\n", l, w, h)
        
    }
	
}