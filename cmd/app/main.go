package main

import (
	"fmt"
	"http-module/internal/greet"
)

func main() {
	msg1 := greet.Hello("Parth Gartan")
	fmt.Println(msg1)
}