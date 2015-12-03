 package main

 import (
         "fmt"
         "os"
         "strings"
 )

 func main() {
         file, err := os.Open("./day1input.txt")
		floor := 0
         if err != nil {
                 fmt.Println(err)
                 os.Exit(1)
         }

         defer file.Close()

// if { // Big loop to iterate through file
		data := make([]byte, 1)
		count, err := file.Read(data)
		fmt.Printf("read %d byte: %q\n", count, data[:count])
		if strings.EqualFold(string(data[:count]), "(") {  // this line doesn't work. bytes can't be compared with strings
			floor = floor + 1
		} else {
			floor = floor - 1
		}
		fmt.Printf("Floor: %d\n", floor)
// }
		
 }
