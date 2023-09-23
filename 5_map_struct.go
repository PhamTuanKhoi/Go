package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	eat   string
	walk  string
	sound int
}

type Fish struct {
	Animal
	loan string `js:"can't havn't"`
}

func MapStruct() {

	// map ->position order change
	member := map[string]string{
		"name": "lee",
	}
	fmt.Println(member)

	people := make(map[string]int)
	people = map[string]int{
		"age": 25,
	}

	people["high"] = 177
	delete(people, "age")

	fmt.Println(people)
	fmt.Println(people["age"])

	_, isExist := people["age"]
	fmt.Println("isExist", isExist)

	// control
	tez := map[string]int{
		"tex": 9,
	}

	tex := tez

	tex["tex"] = 7

	fmt.Println("tex", tez, tex)

	// ========================== struct =================
	dog := Animal{
		eat:   "bone",
		walk:  "four legs",
		sound: 111,
	}

	cat := Animal{}
	cat.eat = "fat"
	cat.walk = "four legs"
	cat.sound = 222

	fmt.Println("struct: ", dog, cat)

	bird := struct {
		eat  string
		walk string
	}{eat: "worm", walk: "two legs"}

	fmt.Println("bird", bird)

	// pointer
	chicken := bird
	chicken.eat = "rice"

	fmt.Println("chicken", chicken)

	// & -> pointer in memory area
	duck := &bird
	duck.eat = "grass"

	fmt.Println("duck", duck, "bird", bird)

	// ========= extent and typeof =======
	fish := Fish{}
	t := reflect.TypeOf(fish)

	field, isExist := t.FieldByName("loan")

	fmt.Println("field", field, isExist)

}
