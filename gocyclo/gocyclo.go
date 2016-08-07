package main

func a(a bool, b bool) {
	if(a || b) {
		if (a && b) {
			print("ab")
		}
		if (a && !b) {n
			print("a")
		}
		if (!a && b) {
			print("b")
		}
	}
}

func main() {
	a(true, false)
}
