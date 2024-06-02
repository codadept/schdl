package util

import (
	"fmt"
	"math"
)

type Int64Func func() (int64, error)

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
