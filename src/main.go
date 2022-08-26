package ungar

import (
	"flag"
	"fmt"
)

func EvaluateHand() {
	var hand string
	var help bool

	flag.StringVar(&hand, "hand", "", "Hand that we want to evaluate")
	flag.BoolVar(&help, "help", false, "Display usage information")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	if hand == "" {
		fmt.Println("Please provide a 7-card hand to evaluate. Use -help for help")
		return
	}

	convertedHand := Convert(hand)
	fmt.Printf("{%d, %d, %d, %d, %d, %d, %d}\n", convertedHand[0], convertedHand[1], convertedHand[2], convertedHand[3], convertedHand[4], convertedHand[5], convertedHand[6])
	value := Evaluate(convertedHand)
	fmt.Printf("Value of your hand is: %v\n", value)
}