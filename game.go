package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var machineHand string
var playerHand string
var handTypes = [3]string{"Rock", "Paper", "Scissors"}

func main() {
	fmt.Println("Let's play Rock, Paper, Scissors!")
	machinePlay()
	playerInput()
	fmt.Println("The machine plays", machineHand)
	fmt.Println("You play", playerHand)
	determineWinner()
}

func machinePlay() {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	machineHand = handTypes[random.Intn(3)]
}

func playerInput() {
	fmt.Println("What will you play?")
	fmt.Println("1 = Rock, 2 = Paper, 3= Scissors")
	var play int     //Variable will be used as int value for rock, paper, or scissors
	var input string //Variable contains user input

	for { //Loop that runs until user gives valid input
		fmt.Scanln(&input)
		if validateInput(input) {
			break
		} else {
			fmt.Println("You can only input 1, 2, or 3")
		}
	}

	play, err := strconv.Atoi(input) //Converts string to int
	if err != nil {
		fmt.Println(err)
	}
	playerHand = handTypes[play-1]
}

func validateInput(input string) bool { //Function checks whether user input is one of three allowed inputs
	if input == "1" || input == "2" || input == "3" {
		return true
	} else {
		return false
	}
}

func determineWinner() { //0 = Rock, 1 = Paper, 2 = Scissors
	switch {
	case machineHand == playerHand:
		fmt.Println("It's a draw")
	case machineHand == handTypes[0] && playerHand == handTypes[1]: // Rock vs Paper
		playerWins()
	case machineHand == handTypes[0] && playerHand == handTypes[2]: // Rock vs Scissors
		machineWins()
	case machineHand == handTypes[1] && playerHand == handTypes[0]: // Paper vs Rock
		machineWins()
	case machineHand == handTypes[1] && playerHand == handTypes[2]: // Paper vs Scissors
		playerWins()
	case machineHand == handTypes[2] && playerHand == handTypes[0]: // Scissors vs Rock
		playerWins()
	case machineHand == handTypes[2] && playerHand == handTypes[1]: // Scissors vs Paper
		machineWins()
	}
}

func playerWins() {
	fmt.Println("You win, well done!")
}

func machineWins() {
	fmt.Println("Machine wins, better luck next time...")
}
