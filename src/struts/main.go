package main

import "fmt"

type Person struct {
	fname string
	lname string
	age   int
}

func main() {

	p := Person{
		fname: "sumit",
		lname: "paul",
		age:   44,
	}

	fmt.Println(p)

}
