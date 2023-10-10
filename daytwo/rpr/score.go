package rpr

import "fmt"

// Opponent Choice
// A = Rock
// B = Paper
// C = Scissors

// My Choice
// X = Rock (Lose)
// Y = Paper (Draw)
// Z = Scissors (Win)

// Result Scores
// Roc = 1
// Pap = 2
// Sci = 3

// Lost = 0
// Draw = 3
// Win  = 6

type Hand string

type ResultScore int

const (
	LostScore ResultScore = 0
	DrawScore ResultScore = 3
	WinScore  ResultScore = 6
)

type Result string

const (
	Lose Result = "X"
	Draw Result = "Y"
	Win  Result = "Z"
)

// const (
// 	Rock     Hand = "ROCK"
// 	Paper    Hand = "PAPER"
// 	Scissors Hand = "SCISSORS"
// )

func (r Result) Score() ResultScore {
	switch r {
	case Lose:
		return LostScore
	case Draw:
		return DrawScore
	case Win:
		return WinScore
	}

	panic("Bad result")
}

type Score struct {
	Score int
}

func NewScore() *Score {
	return &Score{0}
}

func (s *Score) Round(oppoentHand string, resultStr string) {
	op := NewHand(oppoentHand)
	result := Result(resultStr)

	var my Hand
	switch result {
	case Lose:
		my = op.Beats()
		fmt.Println("loses", op, result, result.Score(), my, my.Score())
	case Draw:
		my = op
		fmt.Println("draw", op, result, result.Score(), my, my.Score())
	case Win:
		my = op.Loses()
		fmt.Println("win", op, result, result.Score(), my, my.Score())
	}

	s.Score += int(result.Score()) + my.Score()
}

// Part one
// func (s *Score) Round(oppoentHand string, myHand string) {
// 	op := NewHand(oppoentHand)
// 	my := NewHand(myHand)
//
// 	result := my.Compare(op)
//
// 	s.Score += int(result) + my.Score()
// }

func NewHand(handStr string) Hand {
	switch handStr {
	case "A", "X":
		return Rock()
	case "B", "Y":
		return Paper()
	case "C", "Z":
		return Scissors()
	}

	panic(fmt.Sprintf("Invalid hand string %s", handStr))
}

func Rock() Hand {
	return Hand("Rock")
}

func Paper() Hand {
	return Hand("Paper")
}

func Scissors() Hand {
	return Hand("Scissors")
}

func (h Hand) Compare(other Hand) ResultScore {
	if h == other {
		return DrawScore
	}

	switch h {
	case "Rock":
		if other == "Paper" {
			return LostScore
		} else if other == "Scissors" {
			return WinScore
		}
	case "Paper":
		if other == "Scissors" {
			return LostScore
		} else if other == "Rock" {
			return WinScore
		}
	case "Scissors":
		if other == "Rock" {
			return LostScore
		} else if other == "Paper" {
			return WinScore
		}
	}

	return DrawScore
}

func (h Hand) Beats() Hand {
	switch h {
	case "Rock":
		return Scissors()
	case "Paper":
		return Rock()
	case "Scissors":
		return Paper()
	}

	panic("Bad hand")
}

func (h Hand) Loses() Hand {
	switch h {
	case "Rock":
		return Paper()
	case "Paper":
		return Scissors()
	case "Scissors":
		return Rock()
	}

	panic("Bad hand")
}

func (h Hand) Score() int {
	switch h {
	case "Rock":
		return 1
	case "Paper":
		return 2
	case "Scissors":
		return 3
	default:
		return 0
	}
}
