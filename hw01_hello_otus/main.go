package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	stringOrigin := "Hello, OTUS!"
	fmt.Println("original string", stringOrigin)
	stringReverse := reverse.String(stringOrigin)
	fmt.Println("reverse string", stringReverse)
}
