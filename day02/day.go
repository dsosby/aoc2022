package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	usePart2 := len(os.Args) > 1 && os.Args[1] == "--second"
	if usePart2 {
		Problem2()
	} else {
		Problem1()
	}
}

func Problem1() {
	entries := readEntries()
	total := 0

	for _, play := range entries {
		total += play.getScore()
	}

	fmt.Println(total)
}

func Problem2() {
	// X = lose, Y = draw, Z = win
	entries := readEntries()
	total := 0

	for _, play := range entries {
		total += play.getPlannedScore()
	}

	fmt.Println(total)
}

type Play struct {
	them byte
	you  byte
}

func shapeScore(you rune) int {
	switch you {
	case 'X':
		return 1
	case 'Y':
		return 2
	case 'Z':
		return 3
	}

	return 0
}

// a, x == rock
// b, y == paper
// c, z == scissors

func (p *Play) outcomeScore() int {
	if (p.you == 'X' && p.them == 'C') ||
		(p.you == 'Y' && p.them == 'A') ||
		(p.you == 'Z' && p.them == 'B') {
		return 6 // Win
	}

	if (p.you == 'X' && p.them == 'B') ||
		(p.you == 'Y' && p.them == 'C') ||
		(p.you == 'Z' && p.them == 'A') {
		return 0 // Lose
	}

	return 3 // Draw
}

func (p *Play) getScore() int {
	shapeScore := shapeScore(rune(p.you))
	winScore := p.outcomeScore()
	return shapeScore + winScore
}

// a, x == rock
// b, y == paper
// c, z == scissors
// X = lose, Y = draw, Z = win
func getToPlay(them, strategy byte) byte {
	// Probably could do some maths here
	switch strategy {
	case 'X': //lose
		switch them {
		case 'A':   // rock
			return 'Z' // scissors
		case 'B':
			return 'X'
		case 'C':
			return 'Y'
		}
	case 'Y': // Draw
		switch them {
		case 'A':
			return 'X'
		case 'B':
			return 'Y'
		case 'C':
			return 'Z'
		}
	case 'Z': // Win
		switch them {
		case 'A': // rock
			return 'Y' // paper
		case 'B': // paper
			return 'Z' // scissors
		case 'C': // scissors
			return 'X' // rock
		}
	}

	panic("Bad play")
}

func (p *Play) getPlannedScore() int {
	toPlay := getToPlay(p.them, p.you)

  fmt.Printf("%c,%c is the strategy, you should play %c\n", p.them, p.you, toPlay)

	strategyGame := Play{
		them: p.them,
		you:  toPlay,
	}

	shapeScore := shapeScore(rune(strategyGame.you))
	winScore := strategyGame.outcomeScore()
	return shapeScore + winScore
}

func readEntries() []Play {
	file, fileErr := os.Open("day02/input.txt")
	if fileErr != nil {
		panic(fileErr)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	entries := make([]Play, 0)

	for scanner.Scan() {
		line := scanner.Text()
		entry := Play{
			them: line[0],
			you:  line[2],
		}

		entries = append(entries, entry)
	}

	return entries
}
