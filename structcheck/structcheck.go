package main

import (
	F "fmt"
)

type ie struct {
	a 	string
	c 	string
	d 	bool
	b 	bool
	e 	string
}

func main() {
	s := ie{a: "", b: true, c: ""}
	F.Println(s)
}