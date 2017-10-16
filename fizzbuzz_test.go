package fizzbuzz

import "testing"

func TestFizz(t *testing.T) {

	result := fizz(3)
	if result != "Fizz" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "Fizz")
	}

	result = fizz(5)
	if result != "" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "")
	}

	result = fizz(2)
	if result != "" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "")
	}
}

func TestBuzz(t *testing.T) {

	result := buzz(5)
	if result != "Buzz" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "Buzz")
	}

	result = buzz(3)
	if result != "" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "")
	}
}

func TestNumber(t *testing.T) {

	result := number(2)
	if result != "2" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "2")
	}

	result = number(3)
	if result != "" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "")
	}

	result = number(5)
	if result != "" {
		t.Errorf("result was incorrect, got: %s, want: %s.", result, "")
	}
}
