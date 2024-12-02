package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	totalSafe := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		digits := strings.Fields(line)
		list := []int{}

		for _, digit := range digits {
			intDigit, _ := strconv.Atoi(digit)
			list = append(list, intDigit)
		}

		if isSafeReport(list) || isSafeDampened(list) {
			totalSafe++
		}

	}

	fmt.Println(totalSafe)
}

func isSafeReport(list []int) bool {
	if len(list) < 2 {
		return false
	}

	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(list); i++ {
		diff := math.Abs(float64(list[i] - list[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}
		if list[i] < list[i-1] {
			isIncreasing = false
		}
		if list[i] > list[i-1] {
			isDecreasing = false
		}
	}

	return isIncreasing || isDecreasing
}

func isSafeDampened(list []int) bool {
	for i := 0; i < len(list); i++ {
		newList := append([]int{}, list[:i]...)
		newList = append(newList, list[i+1:]...)
		if isSafeReport(newList) {
			return true
		}
	}
	return false
}
