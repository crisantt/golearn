package main

import (
	"fmt" // This is a standard library format
	"example/greet" // This go file a module.
)

func main() {
	sum := greet.Add(10,10,10)
	fmt.Println(sum)
	mensahe := greet.Hello("Ninja")
	fmt.Println(mensahe)

}
