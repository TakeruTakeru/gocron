package gocron

import (
	"fmt"
	"sort"
)

func min(v []int) int {
	s := append([]int{}, v...)
	sort.Ints(sort.IntSlice(s))
	if len(s) > 0 {
		return s[0]
	}
	return 0
}

func max(v []int) int {
	s := append([]int{}, v...)
	sort.Ints(sort.IntSlice(s))
	l := len(s)
	if l > 0 {
		return s[l-1]
	}
	return 0
}

func getRangeError(fieldName string, from, to int) error {
	if from < to {
		return nil
	}
	return fmt.Errorf("invalid %s range: from -> %d, to -> %d", fieldName, from, to)
}

func getThresholdError(fieldName string, maxAllowed, minAllowed int, values []int) error {
	if max(values) > maxAllowed || minAllowed > min(values) {
		return fmt.Errorf("invalid %s", fieldName)
	}
	return nil
}
