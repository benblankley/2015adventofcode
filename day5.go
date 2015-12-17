package main

import (
	"bufio"
	"fmt"
	"os"
//	"strconv"
	"strings"
)

//--- Day 5: Doesn't He Have Intern-Elves For This? ---
//
//Santa needs help figuring out which strings in his text file are naughty or nice.
//
//A nice string is one with all of the following properties:
//  - It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
//  - It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
//  - It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
//
//For example:
//  - ugknbfddgicrmopn is nice because it has at least three vowels (u...i...o...), a double letter (...dd...), and none of the disallowed substrings.
//  - aaa is nice because it has at least three vowels and a double letter, even though the letters used by different rules overlap.
//  - jchzalrnumimnmhp is naughty because it has no double letter.
//  - haegwjzuvuyypxyu is naughty because it contains the string xy.
//  - dvszwmarrgswjxmb is naughty because it contains only one vowel.
//
// How many strings are nice?


func main() {
	file, err := os.Open("./day5input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println("--- Day 5: Doesn't He Have Intern-Elves For This? ---")
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line is: ", line)
		fmt.Println(strings.ContainsAny(line, "a & e & i & o & u"))
		if (strings.Contains(line, "a") && strings.Contains(line, "e") && strings.Contains(line, "i")) {
			fmt.Println("Contains aei")
		} 
		if (strings.Contains(line, "a") && strings.Contains(line, "e") && strings.Contains(line, "o")) {
			fmt.Println("Contains aeo")
		} 
		if (strings.Contains(line, "a") && strings.Contains(line, "e") && strings.Contains(line, "u")) {
			fmt.Println("Contains aeu")
		} 
		if (strings.Contains(line, "e") && strings.Contains(line, "i") && strings.Contains(line, "o")) {
			fmt.Println("Contains eio")
		}
		if (strings.Contains(line, "e") && strings.Contains(line, "i") && strings.Contains(line, "u")) {
			fmt.Println("Contains eiu")
		}
		if (strings.Contains(line, "i") && strings.Contains(line, "o") && strings.Contains(line, "u")) {
			fmt.Println("Contains iou")
		}

	}
}
