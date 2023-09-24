package main

import "fmt"

func Func() {
	a := 1
	b := 2
	value, err := sum(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)

	p := people{
		name: "lee min ho",
	}

	p.student(a)
}

func sum(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, fmt.Errorf("param not value zero")
	}

	return a + b, nil
}

func (p people) student(a int) {
	fmt.Println("student: ", p, a)
}

type people struct {
	name string
}
