package main

import (
	"fmt"
	"go-oop/person"
)

func main() {
	var p = new(person.Person)
	p.NewPerson("Peter", "Developer")
	fmt.Println(p.ToString())
}
