package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	t, sep2 := "", ""
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	for _, arg := range os.Args[1:] {
		t += sep2 + arg
		sep2 = " "
	}
	fmt.Println("os.Args has a length of", len(os.Args), "First method!\n "+s)
	fmt.Println("second method!\n " + t)
}
