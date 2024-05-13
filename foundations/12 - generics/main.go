package main

type MyInt int

// Adding ~ to the type definition so it doesn't complain the MyInt 😮
type Number interface {
	~int | float64
}

// Generics 😃
func Sum[T Number](m map[string]T) T {
	var sum T

	for _, value := range m {
		sum += value
	}

	return sum
}

// Comparing two values (must be compatible types)
func Compare[T comparable](a T, b T) bool {
	return a == b
}

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	mFloat := map[string]float64{
		"one":   1.1,
		"two":   2.2,
		"three": 3.3,
	}

	mMyInt := map[string]MyInt{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	println(Sum(m))
	println(Sum(mFloat))
	println(Sum(mMyInt))

	println(Compare(1, 1.0))
}
