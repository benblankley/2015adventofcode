package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"math"
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
	robosantaposition := house{0, 0}
	visitedhouses[santaposition]++
	visitedhouses[robosantaposition]++

	var counter float64
	for _, v := range input {
		counter += 1
//		fmt.Printf("%f  %f \n", counter, math.Mod(counter, 2))
		if math.Mod(counter, 2) == 0 {
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
		} else {
                        switch v {
                        case '<':
                                robosantaposition.x--
                        case 'v':
                                robosantaposition.y++
                        case '>':
                                robosantaposition.x++
                        case '^':
                                robosantaposition.y--
                        }
                        visitedhouses[robosantaposition]++
		}
	
	}

	fmt.Printf("Houses visited: %d", len(visitedhouses))


}
