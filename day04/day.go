package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	util "github.com/dsosby/aoc2022/pkg"
)

func main() {
	usePart2 := len(os.Args) > 1 && os.Args[1] == "--second"
	if usePart2 {
		fmt.Println(Problem2())
	} else {
		fmt.Println(Problem1())
	}
}

type Range struct {
	Begin int
	End   int
}

type Assignments struct {
	First  Range
	Second Range
}

func NewAssignments(line string) *Assignments {
	// Simplified -- they are all just ints so pick them out in order
	lineRegexp := regexp.MustCompile(`\d+`)
	intLikes := lineRegexp.FindAll([]byte(line), 4)

	ints := make([]int, 4)
	for i, intLike := range intLikes {
		val, err := strconv.Atoi(string(intLike[:]))
		if err != nil {
			panic(err)
		}

		ints[i] = val
	}

	assignments := Assignments{
		First: Range{
			Begin: ints[0],
			End:   ints[1],
		},
		Second: Range{
			Begin: ints[2],
			End:   ints[3],
		},
	}

	return &assignments
}

// Determines if the other Range is within (inclusive)
// of the parent Range.
func (r *Range) Contains(other *Range) bool {
	return r.Begin <= other.Begin && r.End >= other.End
}

// Determines if the other Range overlaps at all (inclusive)
// with the parent Range.
func (r *Range) Overlaps(other *Range) bool {
  // Lots of comparisons, could short-circuit them or constrain the
  // ranges and use some bitmap technique?
  eitherContains := r.Contains(other) || other.Contains(r)
	startsWithin := other.Begin >= r.Begin && other.Begin <= r.End
	endsWithin := other.End >= r.Begin && other.End <= r.End

	return eitherContains || startsWithin || endsWithin
}

func Problem1() int {
	// Input is pairs of section ranges, inclusive
	// e.g. 2-4,6-8
	pairs := make([]*Assignments, 0, 1)
	for _, line := range util.ReadLines("day04/input.txt") {
		pairs = append(pairs, NewAssignments(line))
	}

	// How many pairs fully contain each other?
	sum := 0
	for _, pair := range pairs {
		if pair.First.Contains(&pair.Second) || pair.Second.Contains(&pair.First) {
			sum += 1
		}
	}

  return sum
}

func Problem2() int {
	// Same as problem 1, but instead of contains use overlaps
	pairs := make([]*Assignments, 0, 1)
	for _, line := range util.ReadLines("day04/input.txt") {
		pairs = append(pairs, NewAssignments(line))
	}

	// How many pairs fully contain each other?
	// Could be more efficiently done as a bitmap, but faster to just refactor
	sum := 0
	for _, pair := range pairs {
		if pair.First.Overlaps(&pair.Second) || pair.Second.Overlaps(&pair.First) {
			sum += 1
		}
	}

  return sum
}
