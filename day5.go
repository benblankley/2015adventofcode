package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	isnice := make(map[string]bool)

	scanner := bufio.NewScanner(file)

	var validID = regexp.MustCompile("(.*[aeiou]){3}")
	var part1, part2, part3 bool

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Printf("%q ", line)

		//		Check for vowel combinations

		if validID.MatchString(line) {
			part1 = true
		//	fmt.Printf(" has 3 vowels ")
		} else {
			part1 = false
		//	fmt.Printf(" does not have 3 vowels ")

		}

		//		Check for double letter combinations
		//		fmt.Println("line contains aa", line, strings.Contains(line, "aa"))

		part2 = false
		var lastChar byte
		for i := 0; i < len(line); i++ {
			if line[i] == lastChar {
				part2 = true
			}
			lastChar = line[i]
		}
		if part2 {
		//	fmt.Printf(" double letter found ")
		} else {
		//	fmt.Printf(" no double letter found ")
		}

		//		Check for special cases
		part3 = true
		if strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy") {
		//	fmt.Printf(" contains ab, cd, pq, or xy ")
			part3 = false
		} else {
		//	fmt.Printf(" no bad chars ")
		}

		// Print some debugging information
		// fmt.Print(part1, part2, part3)
		// fmt.Printf(" | Result: %t \n", isnice[line])
		
		// Calculate nice score for line
		isnice[line] = part1 && part2 && part3

	}

	count := 0
	for _, v := range isnice {
		if v {
			count++
		}
	}

	fmt.Println("Number of nice strings: ", count)
}
