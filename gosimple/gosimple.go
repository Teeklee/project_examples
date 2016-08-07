package main

func main() {
	//Example 1 don't use select with just one case
	var ch chan int
	select {
		case <-ch:
	}

	//Example 2 if can be simplified
	c := true
	if (c == true) {
		//do stuff
	}
	
	//Example 3 Array Copy
	var b1, b2 []byte
	for i, v := range b1 {
		b2[i] = v
	}
}
