package hello

import (
	"fmt"
)

// HelloWorld in a public function because itsname begin by an upppercase
func HelloWorld() string {
	return "Hello world !"
}

// Public class
type Person struct {
	Name string
}

// private function we will test it here
func NewPerson(name string) Person {
	p := Person{Name: name}
	return p
}

func (person Person) SayHello() string {
	return fmt.Sprintf("Hello %s !", person.Name)
}
