package main

import "fmt"

func main() {
	value, isBigger := isBiggerThan10(5, 6)
	fmt.Println(value, isBigger)

	// result, err := buyAlcohol(17)
	result, err := buyAlcohol(21)
	// Error handling
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

// func sum(a, b int) int {
// 	return a + b
// }

func isBiggerThan10(a, b int) (int, bool) {
	if a+b > 10 {
		return a + b, true
	}

	return a + b, false
}

func buyAlcohol(age int) (string, error) {
	if age < 18 {
		return "", fmt.Errorf("You can't buy alcohol")
	}

	return "You can buy alcohol", nil
}
