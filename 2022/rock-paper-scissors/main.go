package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	A = 1
	B = 2
	C = 3
	X = 1
	Y = 2
	Z = 3
)

var shapeVal = map[string]int{
	"A": A,
	"B": B,
	"C": C,
	"X": X,
	"Y": Y,
	"Z": Z,
}

func winningShape(against int) int {
	shapes := []int{1, 2, 3}
	return shapes[against%len(shapes)]
}

func losingShape(against int) int {
	shapes := []int{1, 2, 3}
	return shapes[(against+1)%len(shapes)]
}

func pickWinner(a, b int) int {
	if winningShape(a) == b {
		return b
	}
	return a
}

func play(myShape, againstShape int) int {
	if myShape == againstShape {
		return 3 + myShape
	}
	if pickWinner(myShape, againstShape) == myShape {
		return 6 + myShape
	}
	return myShape
}

func parseGame(s string) (int, int) {
	split := strings.Split(strings.TrimSpace(s), " ")
	a := shapeVal[split[0]]
	b := shapeVal[split[1]]
	return a, b
}

/*
	X - lose
	Y - draw
	Z - win
*/

const (
	Lose = X
	Draw = Y
	Win  = Z
)

func respondTo(shape, desiredOutcome int) int {
	switch desiredOutcome {
	case Draw:
		return shape
	case Win:
		return winningShape(shape)
	case Lose:
		return losingShape(shape)
	}
	panic("invalid shape")
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.TrimSpace(txt) == "" {
			continue
		}
		againstShape, myShape := parseGame(txt)
		score += play(myShape, againstShape)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println("Part one result:", score)
}

func part2() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.TrimSpace(txt) == "" {
			continue
		}
		againstShape, desiredOutcome := parseGame(txt)
		score += play(respondTo(againstShape, desiredOutcome), againstShape)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Println("Part two result:", score)
}

func main() {
	partOne()
	part2()
}
