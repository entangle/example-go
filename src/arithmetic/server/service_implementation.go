package main

import (
	"github.com/entangle/goentangle"
	"definitions/arithmetic"
	"math"
)

// Arithmetic server implementation.
type arithmeticServiceImplementation struct {}

func (s *arithmeticServiceImplementation) DivideInts(a int64, b int64, trace goentangle.Trace) (result int64, err error) {
	// Make sure we won't divide by zero.
	if b == 0 {
		err = arithmetic.DivisionByZeroError.New("division by zero")
		return
	}

	result = a / b
	return
}

func (s *arithmeticServiceImplementation) MultiplyInts(a int64, b int64, trace goentangle.Trace) (result int64, err error) {
	// Quick path out of multiplications by zero.
	if a == 0 || b == 0 {
		result = 0
		return
	}

	// Make sure we won't get overflows.
	overflow := false

	if a < 0 && b > 0 {
		overflow = (a < math.MinInt64 / b)
	} else if a > 0 && b < 0 {
		overflow = (b < math.MinInt64 / a)
	} else if a < 0 {
		overflow = a == math.MinInt64 || b == math.MinInt64 || (-a > math.MaxInt64 / -b)
	} else {
		overflow = (a > math.MaxInt64 / b)
	}

	if overflow {
		err = arithmetic.OverflowError.New("integer overflow")
		return
	}

	result = a * b
	return
}
