package main

import (
	"fmt"
	"strconv"
)

// file different access Number
var Number int = 10

var (
	i float32
	k string = "okay"
)

func Variable() {
	var f float64 = 7.9
	var j int = int(f)
	fmt.Println(j)

	var c string = strconv.Itoa(j)

	fmt.Printf("%v, %T", c, c)

	fmt.Println(k)
	i = 7.7
	fmt.Printf("%v, %T/n", i, i)

}
