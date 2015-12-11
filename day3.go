package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	file, err := ioutil.ReadFile("day3input.txt"])

	// create new struct type called "house", which is basically an array of x and y coordinates
	type house struct {
		x, y int
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	visitedhouses := make(map[house]bool)
	
 		var santa house 
  		for _, count := range file { 
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
 		visitedhouses[*santa] = true
		} 
		
	fmt.Printf("Houses by just santa: %d", len(visitedhouses))
}
