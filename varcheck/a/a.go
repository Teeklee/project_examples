package main

import (
	F "fmt"
)

const c1 = "unused"	//Shadowed
const c2 = "unused"

func main() {
	var c1 = "used"
	F.Println(c1)
}