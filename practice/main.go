package main

import (
	"fmt"
)

func main() {


	type person struct {
		name string 
		age int
		pet string
	}

	bob := person{
		name: "Bob",
		age: 25,
		pet: "cat",
	}

	fmt.Println(bob.name)

}