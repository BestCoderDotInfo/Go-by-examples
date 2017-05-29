package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alex", age: 30})

	fmt.Println(person{name: "Robin"})

	fmt.Println(&person{name: "Ted", age: 25})

	s := person{name: "Barney", age: 27}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 33
	fmt.Println(sp.age)
}
