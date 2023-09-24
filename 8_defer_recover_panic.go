package main

import (
	"fmt"
)

func DeferRecoverPanic() {
	// defer -> task first in last out
	number := 1
	defer fmt.Println("number: ", number)
	fmt.Println("log...")
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	number = 2

	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("Error: ", err)
	// 	}
	// }()

	// panic("chia cho mot so la 0")

	fmt.Println("start")
	panicker()
	fmt.Println("end")

}

func panicker() {
	fmt.Println("panicker")
	// defer -> program next
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error", err)
			// throw error outside
			// panic(err)
		}
	}()

	panic("error server")
}
