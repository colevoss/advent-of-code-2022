package main

import (
	"aoc/daytwo/rpr"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Opponent Choice
// A = Rock
// B = Paper
// C = Scissors

// My Choice
// X = Rock
// Y = Paper
// Z = Scissors

// Result Scores
// Roc = 1
// Pap = 2
// Sci = 3

// Lost = 0
// Draw = 3
// WIn  = 6

func main() {
	// f, err := os.Open("./daytwo/test_input.txt")
	f, err := os.Open("./daytwo/input.txt")

	check("file open", err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score := rpr.NewScore()

	for scanner.Scan() {
		txt := scanner.Text()

		hands := strings.Split(txt, " ")

		// fmt.Println(hands[0])
		// fmt.Println(hands[1])

		score.Round(hands[0], hands[1])
	}

	fmt.Println("Score", score.Score)
}

func check(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
