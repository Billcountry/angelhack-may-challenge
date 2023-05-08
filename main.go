package main

import (
	"fmt"

	"github.com/billcountry/angelhack-may-challenge/question1"
	"github.com/billcountry/angelhack-may-challenge/question2"
)

func run_question1() {
	const instructions = `<<<<<<><><><><<<<><><><><><<<<><><><><><>>>><<><><><><><><><><>>>><<<<
<><><><><><<<<<><><><><><><<<<><><><><><><><><><><><<<<<<><><<><><>>><
<>><<><<>><><<><><><><><><><<<<<<<<<>><<><><<<><><><><<<<<<>>>>>>>>>>>
<>><><><>><<<><><><><<><><<><><><><><><><<<<><><><>><<>>>>><><><>><<<>
<><><><><><>><><><><><><><><><><><><><><><><><<<><><><><><><><><><><><
><><><><><><>>>><><><><><><><><><>><<<<<<<<<<>>>>><<<<<>>>><<<<>><<><<
><><><><><><><><><><<<<<<<><><<><<><<><<><><><><><<>><><>><><><><><<><
<<<<>><<<<><><<<><>>><<><>>>>><>>><<><<><><><><<>><><><><><><><><><><>
<><><><><><<<<><><<<<><<<>>>>>>>>><<><<<>>>>><<<<<<<<<>>>><<><>><><<><
<>><<>><<>><`
	fmt.Printf("Question 1: %d\n", question1.FindFloor(instructions))
}

func run_question2() {
	specs := []*question2.HumanSpecs{
		{Speed: 10, Burst: 6, Rest: 20},
		{Speed: 8, Burst: 8, Rest: 25},
		{Speed: 12, Burst: 5, Rest: 16},
		{Speed: 7, Burst: 7, Rest: 23},
		{Speed: 9, Burst: 4, Rest: 32},
		{Speed: 5, Burst: 9, Rest: 18},
	}
	fmt.Printf("Question 2: %d\n", question2.FarthestDistance(specs, 1234))
}

func main() {
	run_question1()
	run_question2()
}
