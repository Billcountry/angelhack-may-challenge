package question3

import (
	"fmt"
	"sort"
	"strconv"
)

// Only considers permutations divisible by 7
func generatePermutations(numRune []rune, left, right int) []int {
	output := []int{}
	if left == right {
		perm, _ := strconv.ParseInt(string(numRune), 10, 64)
		if (perm % 7) == 0 {
			output = append(output, int(perm))
		}
	} else {
		for i := left; i <= right; i++ {
			numRune[left], numRune[i] = numRune[i], numRune[left]
			output = append(output, generatePermutations(numRune, left+1, right)...)
			numRune[left], numRune[i] = numRune[i], numRune[left]
		}
	}
	return output
}

func IsPermutationDivisibleBy7(num int) float64 {
	numStr := fmt.Sprint(num)
	permutations := generatePermutations([]rune(numStr), 0, len(numStr)-1)

	length := len(permutations)

	if length == 0 {
		return 0.0
	}

	sort.Ints(permutations)

	first := float64(permutations[0])
	last := float64(permutations[length-1])

	return (first + last) / 2.0
}

/*
# Question 3: Is the permutation divisible by 7? (1 mark)

Given an integer string, create all integer permutations of its digits. Determine if there is a permutation whose integer value is evenly divisible by 7, i.e. (permutation value) mod 7 = 0.

For example, the possible permutations of 789 are p = {789, 798, 879, 897, 978, 987}. Of these values, p[2] and p[5] is divisible by 7 because 879 mod 7 = 0 and 987 mod 7 = 0.

Their average is (879 + 987) / 2 = 933

What you’ll need to do is determine if any of the permutations of 1867 are divisible by 7, and if so, what is the average between the smallest and largest permutation? Decimals are allowed.
*/
