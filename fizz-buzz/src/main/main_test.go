package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_should_return_string_for_integer(t *testing.T) {
	assert.Equal(t, "2", fizzbuzz(2))
}

func Test_should_return_fizz_for_three(t *testing.T) {
	assert.Equal(t, "fizz", fizzbuzz(3))
}

func Test_should_return_buzz_for_five(t *testing.T) {
	assert.Equal(t, "buzz", fizzbuzz(20))
}

func Test_should_return_fizz_if_divisible_by_three(t *testing.T) {
	assert.Equal(t, true, isDivisible(9, 3))
	assert.Equal(t, false, isDivisible(8, 3))
	assert.Equal(t, "fizz", fizzbuzz(9))
}

func Test_should_return_fizz_if_divisible_by_five(t *testing.T) {
	assert.Equal(t, true, isDivisible(20, 5))
	assert.Equal(t, false, isDivisible(8, 5))
	assert.Equal(t, "buzz", fizzbuzz(20))
}

func Test_should_return_fizz_if_divisible_by_three_and_five(t *testing.T) {
	assert.Equal(t, true, isDivisible(15, 3))
	assert.Equal(t, true, isDivisible(15, 5))
	assert.Equal(t, false, isDivisible(8, 3))
	assert.Equal(t, false, isDivisible(8, 5))
	assert.Equal(t, "fizzbuzz", fizzbuzz(15))
}
