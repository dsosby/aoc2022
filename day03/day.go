package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  usePart2 := len(os.Args) > 1 && os.Args[1] == "--second"
  if (usePart2) {
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
  fmt.Println("Solution Two")
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
