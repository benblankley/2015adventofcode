package main

import (
	"fmt"
	"io/ioutil"
)

// create new struct type called "house", which is basically an array of x and y coordinates
type house struct {
	x, y int
}

type presentCount map[house]int

func main() {
	inputfile, err := ioutil.ReadFile("day3input.txt")

	if err != nil {
		panic(err)
	}
	visitedhouses := make(presentCount)
	
 		var santa house 
  		for _, count := range inputfile { 
 			switch count { 
 			case '<': 
 				santa.x-- 
 			case 'v': 
 				santa.y++ 
 			case '>': 
 				santa.x++ 
 			case '^': 
 				santa.y-- 
 			} 
 		visitedhouses[santa]++
		} 
		
	fmt.Printf("Houses by Santa: %d", len(visitedhouses))
}
