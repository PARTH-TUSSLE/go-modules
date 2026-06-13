package main

import (
	"fmt"
	"go-modules/internal/greet"
)

func main() {
	msg1 := greet.Hello("Parth Gartan")
	fmt.Println(msg1)
}