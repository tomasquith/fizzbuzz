package fizzbuzz

import "fmt"

var fizz = func(i int) string {
	if i%3 == 0 {
		return "Fizz"
	}
	return ""
}

var buzz = func(i int) string {
	if i%5 == 0 {
		return "Buzz"
	}
	return ""
}

var number = func(i int) string {
	if fizz(i) == "" && buzz(i) == "" {
		return fmt.Sprintf("%d", i)
	}
	return ""
}

func Fizz(i int) string {
	return fizz(i)
}

func Buzz(i int) string {
	return buzz(i)
}

func Number(i int) string {
	return number(i)
}
