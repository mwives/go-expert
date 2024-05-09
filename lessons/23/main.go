package main

func main() {
	a := 1
	b := 2
	c := 3

	if a < b {
		println("a is less than b")
	}

	if a > b && b < c {
		println("a is greater than b and b is less than c")
	}

	if a > b || b < c {
		println("a is greater than b or b is less than c")
	}

	if !(a > b) {
		println("a is not greater than b")
	}
}
