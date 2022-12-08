package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func OpenFile(fname string) *os.File {
	abs_fname, err := filepath.Abs(fname)
	fmt.Println(abs_fname)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(abs_fname)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func MaxInArray(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}

func SumSlice(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
