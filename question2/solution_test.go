package question2_test

import (
	"testing"

	"github.com/billcountry/angelhack-may-challenge/question2"
)

func TestDistanceTraveled(t *testing.T) {
	// John can run 10 m/s for 6 seconds, but then must rest for 20 seconds
	john := question2.HumanSpecs{
		Speed: 10,
		Burst: 6,
		Rest:  20,
	}

	if question2.DistanceTraveled(&john, 100) != 240 {
		t.Fail()
	}

	james := question2.HumanSpecs{
		Speed: 8,
		Burst: 8,
		Rest:  25,
	}
	if question2.DistanceTraveled(&james, 100) != 200 {
		t.Fail()
	}
}

func TestFarthestDistance(t *testing.T) {
	farthest := question2.FarthestDistance([]*question2.HumanSpecs{
		{10, 6, 20},
		{8, 8, 25},
	}, 100)

	if farthest != 240 {
		t.Fail()
	}
}
