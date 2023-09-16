package main

import "fmt"

const (
	_ = 1 << iota * 10
	b
	c
	d
)

func ConstantEnum() {
	const a int = 9
	// a = 8
	fmt.Println(a)

	fmt.Println(b, c, d)
}
