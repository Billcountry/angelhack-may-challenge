package question7_test

import (
	"testing"

	"github.com/billcountry/angelhack-may-challenge/question7"
)

func TestGeneration(t *testing.T) {
	const gen0 = "....XX..X.X..XX..X..X...."
	const gen1 = "X..X.XXXX.XXX.XXX.XX.XX.."
	const gen2 = "XXXXX....X....X...X.X.XXX"

	if question7.Generation(gen0) != gen1 {
		t.Error(question7.Generation(gen0))
		t.Error("Unexpected generation 1 result")
	}

	if question7.Generation(gen1) != gen2 {
		t.Error(question7.Generation(gen1))
		t.Error("Unexpected generation 2 result")
	}
}
