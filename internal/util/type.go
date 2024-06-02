// Package util provides utility functions for the schdl application.
package util

import (
	"fmt"
	"math"
)

// Int64Func represents a function that returns an int64 and an error.
type Int64Func func() (int64, error)

// ConvertInt64ToInt converts the result of a function returning int64 to an int.
// It takes a function fn of type Int64Func as input, calls it to get the int64 result,
// and converts it to int. It returns the converted int result and any error encountered.
func ConvertInt64ToInt(fn Int64Func) (int, error) {
	value, err := fn()
	if err != nil {
		return 0, err
	}

	if value > int64(math.MaxInt) || value < int64(math.MinInt) {
		return 0, fmt.Errorf("value %d out of int range", value)
	}
	return int(value), nil
}
