package main

import "fmt"

func main() {
	x := 2
	p := &x // p now holds the "address if x" and has a type of *int, read as pointer to int
	// *p (pointer to p) holds value of the variable x
	pBefore := x
	*p = 6
	fmt.Println("Value of p before is: ", pBefore, "and value of p after is: ", *p, "and x now has a value of ", x)

	incrP := 1
	pp := incr(&incrP)
	fmt.Println("first value of incrp: ", incrP, "next value is: ", pp, "and third value is: ", incr(&incrP))
}

func incr(p *int) int {
	*p++
	return *p
}
