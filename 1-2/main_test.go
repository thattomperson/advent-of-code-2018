package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		a []int
		r int
	}{
		{[]int{1, -1}, 0},
		{[]int{+3, +3, +4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
	}
	for i, tC := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tC.r, Rept(tC.a...))
		})
	}
}
