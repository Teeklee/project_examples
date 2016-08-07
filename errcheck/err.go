package main

import (
	"fmt"
)

func a() error {
	fmt.Println("this function returns an error")
	return nil
}

func main() {
	_ = a()
	a()
}
