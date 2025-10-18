package main

import (
	"example/greet" // This go file a module.
	"fmt"           // This is a standard library format
)

func main() {
	sum := greet.Add(10, 10, 10)
	fmt.Println(sum)
	mensahe := greet.Hello("Ninja")
	fmt.Println(mensahe)
}
