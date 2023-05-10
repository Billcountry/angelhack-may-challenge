package question4_test

import (
	"testing"

	"github.com/billcountry/angelhack-may-challenge/question4"
)

type testDetails struct {
	efficiencies []int
	expected     int
}

func TestMinimumCost(t *testing.T) {
	tests := []testDetails{
		{[]int{1, 2, 4}, 1},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{4, 2, 8, 1, 9}, 2},
		{[]int{1, 2, 11, 9, 8}, 2},
		{[]int{13, 8, 9, 11, 1}, 3},
	}

	for _, test := range tests {
		res := question4.MinimumCost(test.efficiencies)
		if res != test.expected {
			t.Errorf("%v: %d != %d\n", test.efficiencies, res, test.expected)
		}
	}
}
