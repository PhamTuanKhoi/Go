package main

import (
	"fmt"
)

func IfSwitch() {

	list := map[string]int{
		"one": 1,
		"two": 2,
	}

	if age, isExist := list["one"]; isExist {
		fmt.Println(age)
	}

	number := 7
	switch number {
	case 1:
		fmt.Println(1)
	case 7:
		fmt.Println(7)
	default:
		fmt.Println(-1)
	}

	switch {
	case number == 1:
		fmt.Println(1)
	case number <= 9:
		if true {
			fmt.Println(9)
			break
		}
		fmt.Println(9)
	case number == 7:
		fmt.Println(7)
	default:
		fmt.Println(-1)
	}

	//
	var i interface{} = 1

	switch i.(type) {

	case int:
		fmt.Println("int")

	case string:
		fmt.Println("string")
	}
}
