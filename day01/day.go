package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
  file, fileErr := os.Open("day01/input.txt")
  if fileErr != nil {
    panic(fileErr)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  max := 0
  cur := 0

  for scanner.Scan() {
    line := scanner.Text()

    if line == "" {
      if cur > max {
        max = cur
      }

      cur = 0
    } else {
      num, err := strconv.Atoi(line)
      if err != nil {
        panic(err)
      }

      cur += num
    }
  }

  fmt.Printf("%d\n", max)
}

func Problem2() {
  fmt.Println("Solution Two")
}
