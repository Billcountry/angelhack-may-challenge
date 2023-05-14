package question6_test

import (
	"testing"

	"github.com/billcountry/angelhack-may-challenge/question6"
)

func TestDisconnectNodes(t *testing.T) {
	simple := question6.DisconnectNodes("aabbaa")
	if simple != 2 {
		t.Errorf("Expected 2 but found %d", simple)
	}

	harder := question6.DisconnectNodes("aabbbcccaacccaa")
	if harder != 4 {
		t.Errorf("Expected 4 but found %d", harder)
	}
}
