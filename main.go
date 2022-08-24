package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var machineHand string
var playerHand string
var handTypes = [3]string{"Rock", "Paper", "Scissors"}                                         // Sets the three types of hands that can be played
var winOutcomes = [3][2]string{{"Rock", "Scissors"}, {"Paper", "Rock"}, {"Scissors", "Paper"}} // Determines winning hands
var scores Scores

func main() {
	scores = loadScores() // Load historical scores from scores.json
	fmt.Println("\nWelcome to Rock, Paper, Scissors!")
	startMenu()
}

func startMenu() {
	fmt.Println("\n--- MAIN MENU ---")
	fmt.Println("[1] Start Game")
	fmt.Println("[2] Show Scoreboard")
	fmt.Println("[3] Settings")
	fmt.Println("[4] Quit")

	var choice string
loop: //Set a label for looping back if player gives invalid input
	fmt.Scanln(&choice)
	switch {
	case choice == "1":
		startGame()
	case choice == "2":
		readScores(scores)
		time.Sleep(time.Second)
		startMenu()
	case choice == "3":
		fmt.Println("--- SETTINGS ---")
		fmt.Println("[1] Reset scoreboard")
		fmt.Println("[2] Go back")
		for {
			fmt.Scanln(&choice)
			if choice == "1" {
				scores = resetScoreboard()
				startMenu()
				break
			} else if choice == "2" {
				startMenu()
				break
			} else {
				fmt.Println("You must pick an item from the menu and press Enter")
			}
		}
	case choice == "4":
		quitGame()
	default:
		fmt.Println("You must pick an item from the menu and press Enter")
		goto loop //Go back to loop and lets player give new input
	}
}

func startGame() {
	fmt.Println("Let's play Rock, Paper, Scissors!")
	time.Sleep(time.Second)
	machinePlay()
	playerHand = playerInput()
	fmt.Println("You play", playerHand)
	time.Sleep(time.Second)
	fmt.Println("The machine plays", machineHand)
	time.Sleep(time.Second)
	determineWinner()
	time.Sleep(2 * time.Second)
	playAgain()
}

func machinePlay() { //Function randomizes the hand the machine will play for every round
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	machineHand = handTypes[random.Intn(3)]
}

func playerInput() string {
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
	return handTypes[play-1]
}

func validateInput(input string) bool { //Function checks whether user input is one of three allowed inputs
	if input == "1" || input == "2" || input == "3" {
		return true
	} else {
		return false
	}
}

func determineWinner() { // Determines winner in single player vs machine
	hands := [2]string{playerHand, machineHand}
	switch {
	case machineHand == playerHand:
		scores.Draws += 1 // Adds a draw to scoreboard
		fmt.Println("It's a draw")
	case hands == winOutcomes[0] || hands == winOutcomes[1] || hands == winOutcomes[2]:
		playerWins()
	default:
		machineWins()
	}
}

func determineWinnerPvP(player1 string, player2 string) { // Determines winner in Player vs Player
	hands := [2]string{player1, player2}
	switch {
	case player1 == player2:
		//Draw and play again
	case hands == winOutcomes[0] || hands == winOutcomes[1] || hands == winOutcomes[2]:
		// Player 1 wins
	default:
		// Player 2 wins
	}
}

func playerWins() {
	scores.Wins += 1 // Adds a win to scoreboard
	fmt.Println("You win, well done!")
}

func machineWins() {
	scores.Losses += 1 // Adds a loss to scoreboard
	fmt.Println("Machine wins, better luck next time...")
}

func playAgain() {
	fmt.Println("Do you want to play again?")
	fmt.Println("[1] Yes \n[2] No")
	var response string
loop: //Set a label for looping back if player gives invalid input
	fmt.Scanln(&response)
	if response == "1" {
		startGame()
	} else if response == "2" {
		startMenu()
	} else {
		fmt.Println("You must pick an item from the menu and press Enter")
		goto loop //Go back to loop and lets player give new input
	}
}

func quitGame() {
	updateScoreboard(scores) //Updates scores.json before program end
	fmt.Println("Thank you for playing")
	os.Exit(0)
}
