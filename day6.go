package main

import (
	"bufio"
	"fmt"
	"os"
//	"regexp"
	"strings"
	"strconv"
)

// Because your neighbors keep defeating you in the holiday house decorating 
// contest year after year, you've decided to deploy one million lights in a 
// 1000x1000 grid.
// 
// Furthermore, because you've been especially nice this year, Santa has mailed 
// you instructions on how to display the ideal lighting configuration.
// 
// Lights in your grid are numbered from 0 to 999 in each direction; the lights 
// at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions 
// include whether to turn on, turn off, or toggle various inclusive ranges 
// given as coordinate pairs. Each coordinate pair represents opposite corners 
// of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore 
// refers to 9 lights in a 3x3 square. The lights all start turned off.
// 
// To defeat your neighbors this year, all you have to do is set up your lights 
// by doing the instructions Santa sent you in order.
// 
// For example:
// 
//   - turn on 0,0 through 999,999 would turn on (or leave on) every light.
//   - toggle 0,0 through 999,0 would toggle the first line of 1000 lights, 
//     turning off the ones that were on, and turning on the ones that were 
//     off.
//   - turn off 499,499 through 500,500 would turn off (or leave off) the
//     middle four lights.
// 
// After following the instructions, how many lights are lit?

func main() {
	file, err := os.Open("./day6input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println("--- Day 6: Probably a Fire Hazard ---")

	lights := [1000][1000]bool{}

	scanner := bufio.NewScanner(file)
	x_start := 0
	y_start := 0
	x_end := 0
	y_end := 0

	// Main Loop
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%q ", line)
		
		splitline := strings.Split(line, " ")
		fmt.Printf("%q \n", splitline)

		if splitline[0] == "toggle" {
			// Toggle lights
			s1 := strings.Split(splitline[1], ",")
			x_start, _ = strconv.Atoi(s1[0])
			y_start, _ = strconv.Atoi(s1[1])
			s2 := strings.Split(splitline[3], ",")
			x_end, _ = strconv.Atoi(s2[0])
			y_end, _ = strconv.Atoi(s2[1])
			for i := x_start; i < x_end; i++ {
				for j := y_start; j < y_end; j++ {
				    lights[i][j] = !lights[i][j]
				}
			}
		} else if splitline[0] == "turn" && splitline[1] == "on" {
			// Turn on lights
			s1 := strings.Split(splitline[2], ",")
			x_start, _ = strconv.Atoi(s1[0])
			y_start, _ = strconv.Atoi(s1[1])
			s2 := strings.Split(splitline[4], ",")
			x_end, _ = strconv.Atoi(s2[0])
			y_end, _ = strconv.Atoi(s2[1])
                        for i := x_start; i < x_end; i++ {
                                for j := y_start; j < y_end; j++ {
                                    lights[i][j] = true
                                }
                        }
		} else if splitline[0] == "turn" && splitline[1] == "off" {
			// Turn off lights
			s1 := strings.Split(splitline[2], ",")
                        x_start, _ = strconv.Atoi(s1[0])
                        y_start, _ = strconv.Atoi(s1[1])
			s2 := strings.Split(splitline[4], ",")
                        x_end, _ = strconv.Atoi(s2[0])
                        y_end, _ = strconv.Atoi(s2[1])
                        for i := x_start; i < x_end; i++ {
                                for j := y_start; j < y_end; j++ {
                                    lights[i][j] = false
                                }
                        }
		}
		// fmt.Printf("%d\n", instruction)
		fmt.Printf("x_start %d\n", x_start)
		fmt.Printf("y_start %d\n", y_start)
		fmt.Printf("x_end %d\n", x_end)
		fmt.Printf("y_end %d\n", y_end)

	}
	lightcount := 0
        for i := 0; i < 1000; i++ {
        	for j := 0; j < 1000; j++ {
                     if lights[i][j] { lightcount++ }
                }
      	}

	fmt.Println("Number of lights on: ", lightcount)
}
