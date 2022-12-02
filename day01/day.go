package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func readCalories() []int {
  file, fileErr := os.Open("day01/input.txt")
  if fileErr != nil {
    panic(fileErr)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  calories := make([]int, 10)
  var cur int = 0

  for scanner.Scan() {
    line := scanner.Text()

    if line == "" {
      calories = append(calories, cur)
      cur = 0
    } else {
      num, err := strconv.Atoi(line)
      if err != nil {
        panic(err)
      }

      cur += num
    }
  }

  return calories
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
  calories := readCalories()
  sort.Slice(calories, func(i, j int) bool {
    return calories[i] > calories[j]
  })
  
  top := calories[:3]
  sum := 0
  for _, cal := range top {
    sum += cal
  }

  fmt.Println(sum)
}
