package main

import "fmt"

type Worker interface {
	Work()
}

type Person struct {
	name string
}

func (p Person) Work() {
	fmt.Printf("%s is working", p.name)
}

func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}
func main() {
	p := Person{
		name: "Naveen",
	}

	var w Worker = p
	describe(w)
	w.Work()
}
