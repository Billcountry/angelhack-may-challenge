package question5_test

import (
	"testing"

	"github.com/billcountry/angelhack-may-challenge/question5"
)

func TestDecode(t *testing.T) {
	result := question5.Decode("0110")
	if result != "bc" {
		t.Error("Invalid result", result)
	}
}
