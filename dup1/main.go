package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter some keys to begin with")

	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}
	fmt.Printf("End of the file. and counts has a length of %d  \n", len(counts))

	for line, n := range counts {
		// fmt.Printf("inside the for loop now value of n is: %d \t %s\n", n, line)
		if n > 1 {
			fmt.Printf("\nrepeated files:\n %d \t %s \n", n, line)
		}
	}
}
