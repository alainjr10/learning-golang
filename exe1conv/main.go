package main

import (
	"bufio"
	"fmt"
	"main/exe1conv/convpkg"
	"os"
	"strconv"
)

func main() {
	var inputArgs []string
	if len(os.Args) < 2 {
		fmt.Println("no inputs were passed to command line input")
		fmt.Println("Enter a value to convert")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			}
			inputArgs = append(inputArgs, scanner.Text())
		}
		fmt.Printf("Value of input args is: %v\n", inputArgs)
		computeResults(inputArgs)
	} else {
		inputArgs = os.Args[1:]
		computeResults(inputArgs)
	}

}

func computeResults(inputArgs []string) {
	for _, args := range inputArgs {
		fmt.Printf("args len is: %d OS args is %s\n", len(os.Args), args)
		convertedInput, err := strconv.ParseFloat(args, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing string %v\n", err)
			os.Exit(1)
		}

		cInput := convpkg.Celsius(convertedInput)
		fInput := convpkg.Fahrenheit(convertedInput)

		inputToC := convpkg.FtoC(fInput)
		inputToF := convpkg.CtoF(cInput)

		fmt.Printf("%s = %s: and %s = %s\n", cInput, inputToF, fInput, inputToC)

	}
}
