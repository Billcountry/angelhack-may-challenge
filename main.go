package main

import (
	"fmt"

	"github.com/billcountry/angelhack-may-challenge/question1"
	"github.com/billcountry/angelhack-may-challenge/question2"
	"github.com/billcountry/angelhack-may-challenge/question3"
	"github.com/billcountry/angelhack-may-challenge/question4"
	"github.com/billcountry/angelhack-may-challenge/question5"
	"github.com/billcountry/angelhack-may-challenge/question6"
	"github.com/billcountry/angelhack-may-challenge/question7"
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
	fmt.Printf("- **Question 1:** `%d`\n", question1.FindFloor(instructions))
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
	fmt.Printf("- **Question 2:** `%d`\n", question2.FarthestDistance(specs, 1234))
}

func run_question3() {
	fmt.Printf("- **Question 3:** `%f`\n", question3.IsPermutationDivisibleBy7(1867))
}

func run_question4() {
	efficiencies := []int{1, 3, 54, 712, 52, 904, 113, 12, 135, 32, 31, 56, 23, 79, 611, 123, 677, 232, 997, 101, 111, 123, 2, 7, 24, 57, 99, 45, 666, 42, 104, 129, 554, 335, 876, 35, 15, 93, 13}
	fmt.Printf("- **Question 4:** `%d`\n", question4.MinimumCost(efficiencies))
}

func run_question5() {
	const encoded = "11111011111111110001111111001011111101011111111100110111111111110001001111110100111100110111111100101111010010111111000111111111110001101111110101110011011111111111000110111101001111110010111111001011011111110100111100110111111111110001011101100011111110111111111001110111111111110001111110111111101011111111110001111110111111100111111111110001111011111110111111110100111111111100010011111101001100111111111100011101111111111010111110111111101011111011111101001111001111111111000100111111010011001111111111000111111011111111110001110011111011111110011111110010111110111111000111011111111111000111111110101111011101111111111100011111111101111111010111111110001100111111111100011111111111000111111111110001111111101011110100111111101011111111110001001111110110111111011011010011111110001111111001111111111100011111101111110100111111111100011111111010111101110111111111110001111111011011110111111110000011111110011101"
	fmt.Printf("- **Question 5:** `%s`\n", question5.Decode(encoded))
}

func run_question6() {
	fmt.Printf("- **Question 6:** `%d`\n", question6.DisconnectNodes("kjslaqpwoereeeeewwwefifjdksjdfhjdksdjfkdfdlddkjdjfjfjfjjjjfjffnefhkjgefkgjefkjgkefjekihutrieruhigtefhgbjkkkknbmssdsdsfdvneurghiueor"))
}

func run_question7() {
	fmt.Printf("- **Question 7:** `%d`\n", question7.FiendRepeatingGeneration("XXXX.X....X..X..X.X.XX.XX"))
}

func main() {
	fmt.Println("# AngelHack Monthly Code Challenge")
	fmt.Println("## May 2023 Challenge Statement")
	fmt.Println("### 🎨Theme: “Algorithms: Coding Puzzles”")
	fmt.Println("\n### Results")
	fmt.Println("`$ go run main.go`")
	run_question1()
	run_question2()
	run_question3()
	run_question4()
	run_question5()
	run_question6()
	run_question7()
}
