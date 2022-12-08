package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/rku1999/adventofcode/2022/util"
)

type Hand struct {
	element string
}

func (h Hand) handValue() int {
	if h.element == "A" || h.element == "X" {
		return 1
	} else if h.element == "B" || h.element == "Y" {
		return 2
	} else {
		return 3
	}
}

func (h Hand) CompareHand(other Hand) int {
	if h.element == "X" {
		if other.element == "A" {
			return h.handValue() + 3
		} else if other.element == "B" {
			return h.handValue()
		} else {
			return h.handValue() + 6
		}
	} else if h.element == "Y" {
		if other.element == "A" {
			return h.handValue() + 6
		} else if other.element == "B" {
			return h.handValue() + 3
		} else {
			return h.handValue()
		}
	} else {
		if other.element == "A" {
			return h.handValue()
		} else if other.element == "B" {
			return h.handValue() + 6
		} else {
			return h.handValue() + 3
		}
	}
}

func findMyHand(other Hand, result string) Hand {
	if result == "X" { // Lose
		if other.element == "A" {
			return Hand{"Z"}
		} else if other.element == "B" {
			return Hand{"X"}
		} else {
			return Hand{"Y"}
		}
	} else if result == "Y" { // draw
		if other.element == "A" {
			return Hand{"X"}
		} else if other.element == "B" {
			return Hand{"Y"}
		} else {
			return Hand{"Z"}
		}
	} else { // win
		if other.element == "A" {
			return Hand{"Y"}
		} else if other.element == "B" {
			return Hand{"Z"}
		} else {
			return Hand{"X"}
		}
	}
}

func main() {
	file := util.OpenFile("input.txt")

	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")
		otherHand, myHand := Hand{lineSplit[0]}, Hand{lineSplit[1]}
		score += myHand.CompareHand(otherHand)
	}

	fmt.Println(score) // part 1

	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	score = 0
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		lineSplit := strings.Split(scanner.Text(), " ")
		otherHand, result := Hand{lineSplit[0]}, lineSplit[1]
		score += findMyHand(otherHand, result).CompareHand(otherHand)
	}

	fmt.Println(score) // part 2
}
