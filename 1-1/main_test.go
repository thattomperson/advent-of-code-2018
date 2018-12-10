package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		a []int
		r int
	}{
		{
			[]int{1, 2, 3},
			6,
		},
		{
			[]int{1, 1, 1},
			3,
		},
		{
			[]int{-1, -1, -1},
			-3,
		},
		{
			[]int{-1, 1, -1},
			-1,
		},
		{
			[]int{1, -2, 1},
			0,
		},
	}
	for i, tC := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, tC.r, Sum(tC.a...))
		})
	}
}
