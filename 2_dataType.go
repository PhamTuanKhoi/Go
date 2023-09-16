package main

import (
	"fmt"
)

func Datatype() {
	var a complex64 = 1 + 2i
	var b complex128 = complex(5, 9)

	fmt.Println(a, b)

	var c string = "hello em"

	var d []byte = []byte(c)

	fmt.Printf("%v, %T\n", d, d)

	var e []rune = []rune(c)

	fmt.Printf("%v, %T\n", e, e)

}

/*
Primitive data type

1 boolean
2 numberics
	integer
		signed integer		int8	int16	int32	int64
		unsigned integer	uint8	unint16	unint32	unint64 -> not minus

	float	float32	float64
	complex	complex64	complex128

3 text
	string
	byte -> UTF-8
	rune -> UTF-32


*/
