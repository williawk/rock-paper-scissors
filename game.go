package main

import (
	"fmt"
	"math/rand"
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
	var play int
	fmt.Scanln(&play)
	playerHand = handTypes[play-1]
}

func determineWinner() {
	switch {
	case machineHand == playerHand:
		fmt.Println("It's a draw")
	case machineHand == handTypes[0] && playerHand == handTypes[1]:
		playerWins()
	case machineHand == handTypes[0] && playerHand == handTypes[2]:
		machineWins()
	case machineHand == handTypes[1] && playerHand == handTypes[0]:
		machineWins()
	case machineHand == handTypes[1] && playerHand == handTypes[2]:
		playerWins()
	case machineHand == handTypes[2] && playerHand == handTypes[0]:
		playerWins()
	case machineHand == handTypes[2] && playerHand == handTypes[1]:
		machineWins()
	}
}

func playerWins() {
	fmt.Println("You win, well done!")
}

func machineWins() {
	fmt.Println("Machine wins, better luck next time...")
}
