package question3_test

import (
	"testing"

	"github.com/billcountry/angelhack-may-challenge/question3"
)

func TestIsPermutationDivisibleBy7(t *testing.T) {
	res := question3.IsPermutationDivisibleBy7(789)
	if res != 892.5 {
		t.Errorf("Permutation result %f != 789", res)
	}
}
