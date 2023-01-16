package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnStringForInteger(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz(2))
}

func TestShouldReturnFizzForThree(t *testing.T) {
	assert.Equal(t, "fizz", fizzbuzz(3))
}

func TestShouldReturnBuzzForFive(t *testing.T) {
	assert.Equal(t, "buzz", fizzbuzz(20))
}

func TestShouldReturnFizzIfDivisibleByThree(t *testing.T) {
	assert.Equal(t, true, isDivisible(9, 3))
	assert.Equal(t, false, isDivisible(8, 3))
	assert.Equal(t, "fizz", fizzbuzz(9))
}

func TestShouldReturnFizzIfDivisibleByFive(t *testing.T) {
	assert.Equal(t, true, isDivisible(20, 5))
	assert.Equal(t, false, isDivisible(8, 5))
	assert.Equal(t, "buzz", fizzbuzz(20))
}

func TestShouldReturnFizzIfDivisibleByThree_andFive(t *testing.T) {
	assert.Equal(t, true, isDivisible(15, 3))
	assert.Equal(t, true, isDivisible(15, 5))
	assert.Equal(t, false, isDivisible(8, 3))
	assert.Equal(t, false, isDivisible(8, 5))
	assert.Equal(t, "fizzbuzz", fizzbuzz(15))
}
