package main

import (
	"fmt"
)

func Array() {
	var arr [6]int
	arr[0] = 1

	// & -> points to the memory area arr (static array)
	b := &arr
	b[2] = 5

	fmt.Println(b)
	fmt.Println(cap(arr))
	fmt.Printf("%v, %T\n", arr, arr)

	// dymanic array
	a := []int{7}
	c := a
	var d []int
	d = a
	d[0] = 8

	fmt.Println(a, c, d)

	// sclice

	start := []int{1, 2, 3, 4, 5}
	l1 := start[1:5]
	l1[0] = 9

	fmt.Println(start, l1, cap(l1))
	// start[1 9 3 4 5] l1[9 3 4 5] cap = 4 -> l1[0] === start[1]

	// mask

	ms := make([]int, 0)
	ms = append(ms, 1, 6, 9, 7)
	ms = append(ms, []int{5, 6}...)
	ms = append(ms[2:3], ms[3:4]...) // append 2 param

	fmt.Println("mask: ", ms, len(ms), cap(ms))
}
