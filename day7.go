package main

// --- Day 7: Some Assembly Required ---
// This year, Santa brought little Bobby Tables a set of wires and 
// bitwise logic gates! Unfortunately, little Bobby is a little under the 
// recommended age range, and he needs help assembling the circuit.
// Each wire has an identifier (some lowercase letters) and can carry a 16-bit 
// signal (a number from 0 to 65535). A signal is provided to each wire by a 
// gate, another wire, or some specific value. Each wire can only get a signal
// from one source, but can provide its signal to multiple destinations. A 
// gate provides no signal until all of its inputs have a signal.

// The included instructions booklet describes how to connect the parts 
// together: x AND y -> z means to connect wires x and y to an AND gate, and 
// then connect its output to wire z.

// For example:
//  - 123 -> x means that the signal 123 is provided to wire x.
//  - x AND y -> z means that the bitwise AND of wire x and wire y is 
//    provided to wire z.
//  - p LSHIFT 2 -> q means that the value from wire p is left-shifted by 2 
//    and then provided to wire q.
//  - NOT e -> f means that the bitwise complement of the value from wire e 
//    is provided to wire f.

// Other possible gates include OR (bitwise OR) and RSHIFT (right-shift). If, 
// for some reason, you'd like to emulate the circuit instead, almost all 
// programming languages (for example, C, JavaScript, or Python) provide 
// operators for these gates.

// For example, here is a simple circuit:

// 123 -> x
// 456 -> y
// x AND y -> d
// x OR y -> e
// x LSHIFT 2 -> f
// y RSHIFT 2 -> g
// NOT x -> h
// NOT y -> i

// After it is run, these are the signals on the wires:

// d: 72
// e: 507
// f: 492
// g: 114
// h: 65412
// i: 65079
// x: 123
// y: 456

// In little Bobby's kit's instructions booklet (provided as your puzzle 
// input), what signal is ultimately provided to wire a?

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {
	file, err := os.Open("./day7input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	fmt.Println("--- Day 7: Some Assembly Required ---")

	wires := make(map[string]uint16)
	scanner := bufio.NewScanner(file)

	// Main Loop
	for scanner.Scan() {
		line := scanner.Text()
//		fmt.Printf("%q \n", line)
		formattedwire := strings.Split(line, " -> ")
		fmt.Printf("%q ", formattedwire)
		wires[formattedwire[1]] = 0

		expression := false
		if strings.Contains(formattedwire[0], "AND") {
			fmt.Print(" contains AND \n") 
			expression = true 
			part := strings.Split(formattedwire[0], " AND ")
			part0 := part[0]
			part1 := part[1]
			wires[formattedwire[1]] = wires[part0] & wires[part1]
//			fmt.Print(wires[formattedwire[1]])
			}
		if strings.Contains(formattedwire[0], "OR") {
			fmt.Print(" contains OR \n") 
			expression = true 
			part := strings.Split(formattedwire[0], " OR ")
			part0 := part[0]
			part1 := part[1]
			wires[formattedwire[1]] = wires[part0] | wires[part1]
			}
		if strings.Contains(formattedwire[0], "LSHIFT") {
			fmt.Print(" contains LSHIFT \n") 
			expression = true 
			part := strings.Split(formattedwire[0], " LSHIFT ")
			lshifted := part[0]
			var iterations uint64
			iterations, _ = strconv.ParseUint(part[1], 10, 64)
			wires[formattedwire[1]] = wires[lshifted] << iterations
			}
		if strings.Contains(formattedwire[0], "RSHIFT") {
			fmt.Print(" contains RSHIFT \n") 
			expression = true 
			part := strings.Split(formattedwire[0], " RSHIFT ")
			rshifted := part[0]
			var iterations uint64
			iterations, _ = strconv.ParseUint(part[1], 10, 64)
			wires[formattedwire[1]] = wires[rshifted] >> iterations
			}
		if strings.Contains(formattedwire[0], "NOT") {
			fmt.Print(" contains NOT \n") 
			expression = true
			part := strings.Split(formattedwire[0], "NOT ")
			part1 := part[1]
//			fmt.Print("Part1: ", part1)
			wires[formattedwire[1]] = 65535 - wires[part1]
			}
		if expression == false {
			fmt.Print(" assigning value \n") 
			value, _ := strconv.ParseUint(formattedwire[0], 10, 64)
			variable := formattedwire[1]
			wires[variable] = uint16(value)
			}
		
	}
//	fmt.Print(wires)
	fmt.Print("Value of a:", wires["a"], "\n")
}
