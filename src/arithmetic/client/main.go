package main

import (
	"fmt"
	"definitions/arithmetic"
	"math"
)

func main() {
	client, err := arithmetic.DialArithmeticService("tcp", "127.0.0.1:5555")

	if err != nil {
		panic(err)
	}

	defer client.Close()

	// Let's multiply some numbers!
	for _, numbers := range []struct {
		a int64
		b int64
	} {
		{ 5, 6 },
		{ 2, math.MaxInt64 },
		{ math.MinInt64, math.MaxInt64 },
		{ math.MinInt64, math.MinInt64 },
		{ 0, 0 },
	} {
		result, err := client.MultiplyInts(numbers.a, numbers.b)
		if err != nil {
			fmt.Println(fmt.Sprintf("%d * %d failed: %v", numbers.a, numbers.b, err))
		} else {
			fmt.Println(fmt.Sprintf("%d * %d = %d", numbers.a, numbers.b, result))
		}
	}

	// And then, divide some numbers.
	for _, numbers := range []struct {
		a int64
		b int64
	} {
		{ 6, 5 },
		{ math.MaxInt64, 2 },
		{ math.MinInt64, math.MaxInt64 },
		{ math.MinInt64, math.MinInt64 },
		{ 0, 0 },
	} {
		result, err := client.DivideInts(numbers.a, numbers.b)
		if err != nil {
			fmt.Println(fmt.Sprintf("%d / %d failed: %v", numbers.a, numbers.b, err))
		} else {
			fmt.Println(fmt.Sprintf("%d / %d = %d", numbers.a, numbers.b, result))
		}
	}
}
