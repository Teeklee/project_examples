package main

import (
	F "fmt"
)

type ie struct {
	a 	string
	b 	bool
	c 	string
	d 	bool
	e 	string
}

func main() {
	s := ie{a: "", b: true, c: "", d: false, e: ""}
	F.Println(s)
}