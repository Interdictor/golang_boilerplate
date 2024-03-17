package fizzbuzz

import (
	"testing"
)

func AssertEqual(t *testing.T, result, expected string) {
	if result != expected {
		t.Errorf("got %v but expected %v", result, expected)
	}
}

func Test_NonSpecialNumber(t *testing.T) {
	givenNumber := 1

	result := FizzBuzz(givenNumber)

	expectedResult := "1"
	AssertEqual(t, result, expectedResult)
}

func TestDivisibleByThree(t *testing.T) {
	givenNumber := 3

	result := FizzBuzz(givenNumber)

	expectedResult := "Fizz"
	if result != expectedResult {
		t.Errorf("expected %v but got %v", expectedResult, result)
	}
}

func TestDivisibleByFive(t *testing.T) {
	givenNumber := 5

	result := FizzBuzz(givenNumber)

	expectedResult := "Buzz"
	if result != expectedResult {
		t.Errorf("expected %v but got %v", expectedResult, result)
	}
}

func TestDivisibleByFifteen(t *testing.T) {
	givenNumber := 15

	result := FizzBuzz(givenNumber)

	expectedResult := "FizzBuzz"
	if result != expectedResult {
		t.Errorf("expected %v but got %v", expectedResult, result)
	}
}

func TestDivisibleBySeven(t *testing.T) {
	givenNumber := 7

	result := FizzBuzz(givenNumber)

	expectedResult := "Bazz"
	if result != expectedResult {
		t.Errorf("expected %v but got %v", expectedResult, result)
	}
}

func TestFinalBoss(t *testing.T) {
	givenNumber := 3 * 5 * 7

	result := FizzBuzz(givenNumber)

	expectedResult := "FizzBuzzBazz"
	if result != expectedResult {
		t.Errorf("expected %v but got %v", expectedResult, result)
	}
}
