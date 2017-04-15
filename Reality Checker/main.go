package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func askQuestion(question string) bool {
	fmt.Println(question)
	return readAnswer()
}

func readAnswer() bool {
	for {
		fmt.Print("Enter answer (y/n): ")

		reader := bufio.NewReader(os.Stdin)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if( answer == "y"  ) {
			return true
		}

		if( answer == "n" ) {
			return false
		}
	}
}

func main() {
	fmt.Println("REALITY CHECKER")
	fmt.Println("-----------------")
	fmt.Println("Lately scientists have theorized that we might live in a computer simulation. " +
		"Astrophysicist Neil deGrasse Tyson have put the odds at 50-50 that our entire existence " +
		"is a program on someone else’s hard drive.")
	fmt.Println()
	fmt.Println("This application will help you to find out if you're living in a computer simulation, or in the real world. " +
		"Please answer the following questions to get your answer.")
	fmt.Println()

	goForward := askQuestion("Do you wish to continue?")
	fmt.Println()
	if( ! goForward ) {
		os.Exit(0)
	}

	dress := askQuestion("Have you met a girl in a red dress lately?")
	fmt.Println()

	pill := askQuestion("Have a stranger offered you a blue and a red pill?")
	fmt.Println()

	var viagra bool

	if(pill) {
		viagra = askQuestion("Was the blue pill a Viagra?")
		fmt.Println()
	}

	dejavu := askQuestion("Do you often experience déjà vu's?")
	fmt.Println()

	fmt.Println("ANSWER")
	fmt.Println("--------")

	if( dress && pill && dejavu && ! viagra ) {
		fmt.Println("You are probably living in a computer *simulation*.")
		fmt.Println()

		fmt.Println("ARE THERE ANY PURPOSE TO LIFE?")
		fmt.Println("Your main purpuse is to create electric power for our robot overlords.")
		fmt.Println()

		fmt.Println("DO I GET AN AFTERLIFE?")
		fmt.Println("Your memory will probably get garbage collected.")
		fmt.Println()

	} else {
		fmt.Println("You are probably *flesh and blood*.")
		fmt.Println()
		fmt.Println("Live, breathe, eat and continue doing human stuff.")
		fmt.Println()
	}
}
