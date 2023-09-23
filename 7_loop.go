package main

import (
	"fmt"
)

func Loop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	j := 1
	for j < 10 {
		fmt.Println("j: ", j)
		j++
	}

	k := 0
	for {
		k++

		if k%2 == 0 {
			continue
		}

		fmt.Println("ojt: ", k)

		if k >= 10 {
			break
		}

	}

	//
Loop:
	for a := 1; a < 5; {
		a++
		fmt.Println("a: ", a)
		for b := 1; b < 5; b++ {
			if a > 1 {
				break Loop
			}
			fmt.Println("a + b ", a+b)
		}
	}

	// arr
	list := []string{"men", "woment"}

	for index, value := range list {
		fmt.Println(index, ": ", value)
	}

	// map
	ages := map[string]int{
		"men":    100,
		"woment": 99,
	}

	for key, value := range ages {
		fmt.Println(key, ": ", value)
	}
}
