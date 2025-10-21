package main

import (
	"fmt"

	"example/greet" // This go file a module.
)

func main() {
	sum := greet.Add(10, 10, 10)
	fmt.Println(sum)
	mensahe := greet.Hello("Ninja")
	fmt.Println(mensahe)
	var aa string
	fmt.Print("Ilagay ang iyong pangalan: ")
	fmt.Scan(&aa, "\n")
}
