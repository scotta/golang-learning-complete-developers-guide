package main

import "fmt"

func ex1() {
	t := 20
	var f float64 = float64(t)
	fmt.Printf("Exercise 1: t=%v f=%v\n", t, f)
}

func ex2() {
	// use an untyped constant; behaves like a literal
	const value = 99
	var t int = value
	var f float64 = value

	fmt.Printf("Exercise 2: t=%v f=%v\n", t, f)
}

func ex3() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Printf("Exercise 3 max value:\tb=%v\tsmallI=%v\tbigI=%v\n", b, smallI, bigI)

	b++
	smallI++
	bigI++

	fmt.Printf("Exercise 3 max value+1:\tb=%v\tsmallI=%v\tbigI=%v\n", b, smallI, bigI)
}

func main() {
	ex1()
	ex2()
	ex3()
}
