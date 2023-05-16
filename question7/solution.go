package question7

import (
	"math"
	"strings"
)

/*
1. At each minute, lifeforms live and die based on the number of lifeforms in the four
adjacent tiles. For reference, this is what the adjacent tiles look like (.) based on one lifeform
(X)
.
.X.
.
2. A lifeform passes away, and is replaced by an empty space UNLESS there is EXACTLY
one lifeform adjacent to it.
3. An empty space becomes filled with a lifeform if exactly ONE or TWO lifeforms are
adjacent to it. Otherwise, the lifeform or the empty space remains the same

HQ also says to calculate the lifeform score of that layout. Due to some odd configuration in
your communicator, you’ll have to calculate it this way:
2**(the tile number)

Question
Consider the following start state:
XXXX.
X....
X..X.
.X.X.
XX.XX

What is the lifeform score for the first layout that appears twice?
*/

func Generation(parents string) string {
	parents_rune := []rune(parents)
	out := []string{}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			current := string(parents_rune[i*5+j])
			top := (i-1)*5 + j
			bottom := (i+1)*5 + j
			left := i*5 + j - 1
			right := i*5 + j + 1
			adjacent_live_forms := 0
			if i > 0 && parents_rune[top] == 'X' {
				adjacent_live_forms++
			}
			if i < 4 && parents_rune[bottom] == 'X' {
				adjacent_live_forms++
			}
			if j > 0 && parents_rune[left] == 'X' {
				adjacent_live_forms++
			}
			if j < 4 && parents_rune[right] == 'X' {
				adjacent_live_forms++
			}

			if current == "X" {
				// A lifeform passes away, and is replaced by an empty space UNLESS there is EXACTLY one lifeform adjacent to it.
				if adjacent_live_forms == 1 {
					out = append(out, "X")
				} else {
					out = append(out, ".")
				}
			} else {
				// An empty space becomes filled with a lifeform if exactly ONE or TWO lifeforms are adjacent to it. Otherwise, the lifeform or the empty space remains the same
				if adjacent_live_forms == 1 || adjacent_live_forms == 2 {
					out = append(out, "X")
				} else {
					out = append(out, ".")
				}
			}
		}
	}
	return strings.Join(out, "")
}

func FiendRepeatingGeneration(gen0 string) int64 {
	generations := []string{gen0}
	last_gen := gen0

	for {
		next_gen := Generation(last_gen)
		for _, gen := range generations {
			if gen == next_gen {
				var score int64 = 0
				for i, state := range []rune(next_gen) {
					if state == 'X' {
						score += int64(math.Pow(2, float64(i)))
					}
				}
				return score
			}
		}
		generations = append(generations, next_gen)
		last_gen = next_gen
	}
}
