package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func maxInArray(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func sumSlice(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	abs_fname, err := filepath.Abs("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(abs_fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

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

	fmt.Println(maxInArray(elfCalories)) // part 1

	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] > elfCalories[j]
	})

	fmt.Println(sumSlice(elfCalories[:3])) // part 2

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
