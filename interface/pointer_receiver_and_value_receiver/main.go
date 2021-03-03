package main

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
	fmt.Printf("Type = %T, value = %v\n", p, p)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("State %s Country %s", a.state, a.country)
	fmt.Printf("Type = %T, value = %v\n", a, a)

}

func describeA(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {
	var d1 Describer
	p1 := Person{
		name: "Sam",
		age:  25,
	}

	d1 = p1
	d1.Describe()
	describeA(d1)

	p2 := Person{
		name: "Jame",
		age:  30,
	}

	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{
		state:   "Washington",
		country: "USA",
	}
	//d2 = a
	d2 = &a
	d2.Describe()
	describeA(d1)
}
