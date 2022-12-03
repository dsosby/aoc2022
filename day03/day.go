package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/slices"
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
	lines := readLines()

	var priorities []int
	for _, line := range lines {
		// Split each line in two
		compartmentSize := len(line) / 2
		left := line[:compartmentSize]
		right := line[compartmentSize:]

		// find the common items between both sides
		common := make(map[rune]bool)
		found := make([]rune, 0)
		for _, rightChar := range right {
			if strings.ContainsRune(left, rightChar) && !common[rightChar] {
				common[rightChar] = true
				found = append(found, rightChar)
			}
		}

		// Calculate the priority of the common items (a = 1)
		// a = 97, A = 64
		for commonRune := range common {
			var priority int

			if commonRune >= 'A' && commonRune <= 'Z' {
				priority = int(commonRune) - 65 + 27
			} else if commonRune >= 'a' && commonRune <= 'z' {
				priority = int(commonRune) - 97 + 1
			}

			priorities = append(priorities, priority)
		}

		fmt.Printf("%v,%v\t%v\t%v\n", left, right, common, priorities)
	}

	// Sum the priorities
	var sum int
	for _, priority := range priorities {
		sum += priority
	}

	fmt.Println(sum)
}

func Problem2() {
	// Group by 3 lines
	lines := readLines()
	groupedLines := chunkBy(lines, 3)

	fmt.Println(groupedLines)

	// Find the common rune in all 3 lines
	groupBadges := make([]rune, 0)
	for _, group := range groupedLines {
		commonFirstTwo := common([]rune(group[0]), []rune(group[1]))
		commonLast := common(commonFirstTwo, []rune(group[2]))
		groupBadges = append(groupBadges, commonLast[0])
	}

	// Calculate the priority of each common rune
	// Sum them up
	sum := 0
	for _, groupBadge := range groupBadges {
		sum += priority(groupBadge)
	}

	fmt.Println(sum)
}

func common(left, right []rune) []rune {
	// find the common items between both sides
	common := make(map[rune]bool)
	found := make([]rune, 0) // Iterate keys of common? lol -- go doesn't do that do it yourself

	for _, rightRune := range right {
		if slices.Contains(left, rightRune) && !common[rightRune] {
			common[rightRune] = true
			found = append(found, rightRune)
		}
	}

	return found
}

func priority(commonRune rune) int {
	pri := 0

	if commonRune >= 'A' && commonRune <= 'Z' {
		pri = int(commonRune) - 65 + 27
	} else if commonRune >= 'a' && commonRune <= 'z' {
		pri = int(commonRune) - 97 + 1
	}

	return pri
}

// Stolen from SO - https://stackoverflow.com/questions/35179656/slice-chunking-in-go
func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func readLines() []string {
	file, fileErr := os.Open("day03/input.txt")
	if fileErr != nil {
		panic(fileErr)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
