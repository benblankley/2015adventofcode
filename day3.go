package main

import (
	"fmt"
	"os"
	"io/ioutil"
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

	inputfile, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		panic(err)
	}
	visitedhouses := make(presentCount)
	
 		var santa house
 		visitedhouses[santa]++
  		for _, vm := range inputfile { 
 			switch vm { 
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
		
	fmt.Println("Houses by Santa: %d", len(visitedhouses))
}
