// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of multiplying them
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of division
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Can't divide by 0")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Number cannot be negative")
	}
	return math.Sqrt(a), nil
}
