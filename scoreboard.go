package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Scores struct {
	Wins   int
	Draws  int
	Losses int
}

func loadScores() Scores {
	content, err := os.ReadFile("scores.json")
	if err != nil {
		fmt.Println(err)
	}
	var currentScores Scores
	err = json.Unmarshal(content, &currentScores)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Scores have been loaded!")
	return currentScores
}

func readScores(currentScores Scores) {
	fmt.Printf("Wins: %d, Draws: %d, Losses: %d\n", currentScores.Wins, currentScores.Draws, currentScores.Losses)
}

func updateScoreboard(updatedScores Scores) {
	content, err := json.Marshal(updatedScores)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("scores.json", content, 0666)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Scores have been updated!")
}

func resetScoreboard() {
	resetScores := Scores{0, 0, 0}
	content, err := json.Marshal(resetScores)
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("scores.json", content, 0666)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Scores have been reset!")
}
