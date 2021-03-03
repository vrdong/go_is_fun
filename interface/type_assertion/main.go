package main

import "fmt"

func assert(i interface{}) {
	v, ok := i.(int)
	if ok {
		fmt.Println(v)
	}
	//fmt.Println(v)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	findType("Cao Ba Dong")
	findType(123)
	findType(69.96)
}
