package problem2

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var oponent = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
}

var player = map[string]string{
	"X": "R",
	"Y": "P",
	"Z": "S",
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
var gameResult = map[string]string{
	"X": "L",
	"Y": "D",
	"Z": "W",
}

var playerScore = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var oponentScore = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

func Strategy(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		score += calculateScore(oponent[line[0]], player[line[1]])
		score += playerScore[line[1]]
	}

	return score
}

func CalculateStrategy(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to open file.")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		play := calculatePlay(oponent[line[0]], gameResult[line[1]])
		score += calculateScore(oponent[line[0]], play)
		score += oponentScore[play]
	}

	return score
}

func calculatePlay(play1, result string) string {
	if result == "D" {
		return play1
	}
	// Loose
	if result == "L" {
		switch play1 {
		case "R":
			return "S"
		case "S":
			return "P"
		case "P":
			return "R"
		}
	}
	// Win
	switch play1 {
	case "R":
		return "P"
	case "S":
		return "R"
	case "P":
		return "S"
	}
	return ""
}

func calculateScore(play1, play2 string) int {
	if strings.Compare(play1, play2) == 0 {
		return 3
	}
	//Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock
	if (play1 == "R" && play2 == "S") || (play1 == "S" && play2 == "P") || (play1 == "P" && play2 == "R") {
		return 0
	}
	return 6
}
