package utils

import (
	"fmt"
)

// ForEachRange is simple utility calling given function for each sub-range
// defined within [0-totalSize]. It wont be called however if range is empty.
// For example, MakeRanges(35, 10, fc) will result in fc being called 4 times with arguments:
// * 0-10
// * 10-20
// * 20-30
// * 30-35
// It can be used to split slice into smaller ones.
// Params totalSize must not be less than 0, maxSubRangeSize must be bigger than 0. If those
// are not met, function panics.
// If given function returns error, iteration stops and ForEachRange returns that error
func ForEachRange(totalSize int, maxSubRangeSize int, handler func(from, to int) error) error {
	if maxSubRangeSize < 1 || totalSize < 0 {
		panic(fmt.Errorf("maxSubRangeSize must be bigger than 0 and totalSize not less than 0: for %d and %d",
			maxSubRangeSize, totalSize))
	}
	from := 0
	for {
		to := from + maxSubRangeSize
		if to > totalSize {
			to = totalSize
		}
		if from == to {
			return nil
		}
		if err := handler(from, to); err != nil {
			return err
		}
		from = to
	}
}
