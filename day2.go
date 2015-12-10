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
	
	area := 0
	scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Printf("line is %[1]T %[1]q ", line)
        coordinates := strings.Split(line, "x")

        l, _ := strconv.Atoi(coordinates[0])
        w, _ := strconv.Atoi(coordinates[1])
        h, _ := strconv.Atoi(coordinates[2])
        fmt.Printf("%d %d %d  cumulative area:%d ", l, w, h, area)
	
	if l*w < w*h && l*w < h*l {
		areanew := (2*l*w + 2*w*h + 2*h*l)
		areaextra := l*w
        fmt.Printf("l*w is smallest\n")
        fmt.Printf("%d %d\n", areanew, areaextra)
	area = area + areanew + areaextra        
	} else if w*h < l*w && w*h < h*l {
		areanew := (2*l*w + 2*w*h + 2*h*l)
		areaextra := w*h
        fmt.Printf("w*h is smallest\n")
        fmt.Printf("%d %d\n", areanew, areaextra)
	area = area + areanew + areaextra        
	} else {
		areanew := (2*l*w + 2*w*h + 2*h*l)
		areaextra := h*l
        fmt.Printf("h*l is smallest\n")
        fmt.Printf("%d %d\n", areanew, areaextra)
	area = area + areanew + areaextra        
	}
    }
        fmt.Printf("Area:  %d\n", area)
	
}
