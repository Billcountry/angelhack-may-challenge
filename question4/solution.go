package question4

import (
	"sort"
)

func MinimumCost(efficiencies []int) int {
	length := len(efficiencies)
	if (length % 2) == 1 {
		lowest := -1
		// Find the smallest possibility by removing one digit at a time
		for i := 0; i < length; i++ {
			possible_slice := []int{}
			possible_slice = append(possible_slice, efficiencies[:i]...)
			possible_slice = append(possible_slice, efficiencies[(i+1):]...)
			res := MinimumCost(possible_slice)
			if lowest < 0 || res < lowest {
				lowest = res
			}
		}
		return lowest
	}

	sort.Ints(efficiencies)

	output := 0
	for i := 0; i < (length / 2); i++ {
		current := i * 2
		output += efficiencies[current+1] - efficiencies[current]
	}
	return output
}

/*
# Question 4: Find the group’s efficiency (3 marks)

A group of workers gathered to complete a task. Each worker has an efficiency rating. They will be grouped in pairs so an even number of workers are required. The cost of a pair is the absolute difference of the efficiencies assigned to the workers. The cost of the task is the sum of the costs of all pairs formed. There are an odd number of workers to choose from, so one worker will not be paired. Select the worker to exclude so the task's cost is minimized.

Given n workers and efficiency for each worker, find a configuration of the workers such that the cost of the task is the minimum possible. Return the minimum cost as the answer.

### First Example
Efficiency = [1, 2, 4]

The first worker has an efficiency of 1, the second worker has an efficiency of 2, and the last worker has an efficiency of 4.

If we excluded the first worker (1) and pair the second (2) and last worker (4), the cost is |2 - 4| = 2
If we excluded the second worker (2), and pair the first and last worker instead, the cost is |1 - 4| = 3
If we excluded the last worker (4), the cost is |1 - 2| = 1

Thus, the minimum cost is 1.

### Second Example
efficiency = [4, 2, 8, 1, 9]

The first worker has an efficiency of 4, the second worker has an efficiency of 2, … , last worker has an efficiency of 9.

If we excluded the first worker (4), and you pair up the second and third workers (2, 8), and the fourth and fifth workers (1, 9), the cost of the task is |2 - 8| + |1 - 9| = 14.

Can we do it differently? If we pair up the second and fourth workers (2, 1) and third and last worker (8, 9), the cost of the task is |2 - 1| + |8 - 9| = 2.

Suppose we exclude the last worker who has an efficiency of 9, because we think that the largest inefficiency is bad. In that case, we have (4, 2, 8, 1), and you won’t be able to get a cost that’s lower than 2.

This is the minimum possible cost so return 2.

## Task
What is the minimum cost of this array of efficiencies:
```[1, 3, 54, 712, 52, 904, 113, 12, 135, 32, 31, 56, 23, 79, 611, 123, 677, 232, 997, 101, 111, 123, 2, 7, 24, 57, 99, 45, 666, 42, 104, 129, 554, 335, 876, 35, 15, 93, 13]```
*/
