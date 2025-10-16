package greet

import "fmt"

func Hello(name string) string{
	message := fmt.Sprintf("Wazzup %v. Welcome sa code ko diha", name)
	return message
}

func Add(x, y, z int) int{
	return x + y + z
}
