package main

import "fmt"

/*
	syntax:
		* -> point memory
		& -> address memory
*/

func Pointer() {
	// ============================= primitive type ================================
	var a int = 3
	// point b to memory of a
	var b *int = &a
	fmt.Println(a, *b)

	a = 5
	fmt.Println(a, *b)

	*b = 7
	fmt.Println(a, *b)

	// ============================= reference type ================================

	// ==============> static array point to memory
	var list []int = []int{1, 2, 3}

	var arr []int = list

	arr[0] = 7

	list[1] = 9

	fmt.Println("list, arr: ", list, arr)

	// ==============> dymanic array NOT point to memory
	array := [3]int{1, 2, 3}

	arrayList := array
	array[0] = 5

	fmt.Println("array, arrayList: ", array, arrayList)

	// ==============> map
	var student map[string]string = map[string]string{"name": "ho"}

	var teacher map[string]string = student

	teacher["name"] = "lee"

	fmt.Println("student, teacher: ", student, teacher)

	// ==============> struct

	var k *AgeStruct
	k = &AgeStruct{age: 7}
	fmt.Println(k) // <nil> === null
	// k = new(AgeStruct)

	(*k).age = 20
	fmt.Println("struct: ", k.age)

}

type AgeStruct struct {
	age int
}
