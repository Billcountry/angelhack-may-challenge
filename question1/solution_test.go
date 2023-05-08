package question1_test

import (
	"testing"

	"github.com/billcountry/angelhack-may-challenge/question1"
)

func TestFindFloor(t *testing.T) {
	if question1.FindFloor("<<>>") != 0 {
		t.Fail()
	}

	if question1.FindFloor("<><>") != 0 {
		t.Fail()
	}

	if question1.FindFloor("<<<") != 3 {
		t.Fail()
	}

	if question1.FindFloor(">><<<>>") != -1 {
		t.Fail()
	}
}
