package main

import (
	"fmt"
	"github.com/thazelart/go-example/internal/hello"
)

func main() {
	// init a var with :=
	pierre := hello.NewPerson("Pier")

	// when already initialized use =
	pierre = hello.NewPerson("Pierre")

	fmt.Println(hello.HelloWorld())
	fmt.Println(pierre.SayHello())
}
