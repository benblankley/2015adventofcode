package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./day2input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	area := 0
	ribbon := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("line is %[1]T %[1]q ", line)
		coordinates := strings.Split(line, "x")

		l, _ := strconv.Atoi(coordinates[0])
		w, _ := strconv.Atoi(coordinates[1])
		h, _ := strconv.Atoi(coordinates[2])
		fmt.Printf("L: %3d W: %3d H:%3d ", l, w, h)
		areanew := (2*l*w + 2*w*h + 2*h*l)

		areaextra := 0
		if l*w <= w*h && l*w <= h*l {
			areaextra = l * w
		} else if w*h <= l*w && w*h <= h*l {
			areaextra = w * h
		} else {
			areaextra = h * l
		}
		area = area + areanew + areaextra
		fmt.Printf("Added paper: %7d Total paper: %7d ", areanew+areaextra, area)

		ribbonnew := 0
		if 2*l+2*w <= 2*w+2*h && 2*l+2*w <= 2*h+2*l {
			ribbonnew = (2*l + 2*w) + (l * w * h)
		} else if 2*w+2*h <= 2*l+2*w && 2*w+2*h <= 2*h+2*l {
			ribbonnew = (2*w + 2*h) + (l * w * h)
		} else {
			ribbonnew = (2*h + 2*l) + (l * w * h)
		}
		ribbon = ribbon + ribbonnew
		fmt.Printf("Added ribbon: %7d Total ribbon: %7d\n", ribbonnew, ribbon)

	}
	fmt.Printf("Total Area:  %d\n", area)
	fmt.Printf("Total Ribbon:  %d\n", ribbon)

}
