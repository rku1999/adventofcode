package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/rku1999/adventofcode/2022/util"
)

func main() {
	file := util.OpenFile("input.txt")

	elfCalories := []int{}
	subSum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		calories := scanner.Text()
		if calories == "" {
			elfCalories = append(elfCalories, subSum)
			subSum = 0
		} else {
			if cal, err := strconv.Atoi(calories); err == nil {
				subSum += cal
			} else {
				fmt.Println("Failed to parse calories to int")
			}
		}
	}

	fmt.Println(util.MaxInArray(elfCalories)) // part 1

	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] > elfCalories[j]
	})

	fmt.Println(util.SumSlice(elfCalories[:3])) // part 2
}
