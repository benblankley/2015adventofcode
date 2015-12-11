package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// create new struct type called "house", which is basically an array of x and y coordinates
type house struct {
	x, y int
}

type presentCount map[house]int

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	visitedhouses := make(presentCount)

	santaposition := house{0, 0}
	visitedhouses[santaposition]++
	for _, v := range input {
		switch v {
		case '<':
			santaposition.x--
		case 'v':
			santaposition.y++
		case '>':
			santaposition.x++
		case '^':
			santaposition.y--
		}
		visitedhouses[santaposition]++
	}
	result := 0
	for coords, prescount := range visitedhouses {
		fmt.Printf("  %s has %d presents\n", coords, prescount)
		result++
	}
	fmt.Printf("%d houses have at least one present\n\n", result)


}
