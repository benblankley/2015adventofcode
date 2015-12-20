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
// -- Part Two ---
// 
// Realizing the error of his ways, Santa has switched to a better model of 
// determining whether a string is naughty or nice. None of the old rules 
// apply, as they are all clearly ridiculous.
// 
// Now, a nice string is one with all of the following properties:
// 
//     It contains a pair of any two letters that appears at least twice in the 
// string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like 
// aaa (aa, but it overlaps).
//     It contains at least one letter which repeats with exactly one letter 
// between them, like xyx, abcdefeghi (efe), or even aaa.
// 
// For example:
// 
//     qjhvhtzxzqqjkmpb is nice because is has a pair that appears twice (qj) 
// and a letter that repeats with exactly one letter between them (zxz).
//     xxyxx is nice because it has a pair that appears twice and a letter that 
// repeats with one between, even though the letters used by each rule overlap.
//     uurcxstgmygtbstg is naughty because it has a pair (tg) but no repeat 
// with a single letter between them.
//     ieodomkazucvgmuy is naughty because it has a repeating letter with one 
// between (odo), but no pair that appears twice.
// 
// How many strings are nice under these new rules?

func main() {
	file, err := os.Open("./day5input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println("--- Day 5: Doesn't He Have Intern-Elves For This? ---")

	isnice := make(map[string]bool)
	// Part 2
	isnice2 := make(map[string]bool)

	scanner := bufio.NewScanner(file)

	var validID = regexp.MustCompile("(.*[aeiou]){3}")
	var part1, part2, part3 bool
	// Part 2
	var part4, part5 bool

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

		part4=false
                for i := 0; i < len(line)-2; i++ {
			xy := line[i:i+2]
			if strings.Contains(line[i+2:], xy) {
          	              part4 = true
		              //fmt.Printf(" double doubles on line: %q \n", line)
                        }
                }

		part5 = false
                for i := 0; i < len(line)-2; i++ {
                        if line[i] == line[i+2] {
                                part5 = true
                                //fmt.Printf(" xyx on line: %q \n", line)
                        }
                }


		// Print some debugging information
		// fmt.Print(part1, part2, part3)
		// fmt.Printf(" | Result: %t \n", isnice[line])
		
		// Calculate nice score for line
		isnice[line] = part1 && part2 && part3
		isnice2[line] = part4 && part5
	}

	count := 0
	for _, v := range isnice {
		if v {
			count++
		}
	}

	// Part 2
	count2 := 0
	for _, v := range isnice2 {
		if v {
			count2++
		}
	}


	fmt.Println("Number of nice strings (Part 1): ", count)
	fmt.Println("Number of nice strings (Part 2): ", count2)
}
