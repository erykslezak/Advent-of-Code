package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Opens file, logs error if any, closes file at end of main().
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Two empty slices.
	list1 := []int{}
	list2 := []int{}

	// Create a new scanner to read the file line by line.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Pulls each line from text file and retrieves digits from line by splitting fields.
		line := scanner.Text()
		digits := strings.Fields(line)

		// Converts from string to int and assigns digits to variables, no error check.
		firstNum, _ := strconv.Atoi(digits[0])
		secondNum, _ := strconv.Atoi(digits[1])

		// Appends each digit to corresponding slice.
		list1 = append(list1, firstNum)
		list2 = append(list2, secondNum)
	}
	// Sorts slices from lowest to highest.
	sort.Slice(list1, func(i, j int) bool { return list1[i] < list1[j] })
	sort.Slice(list2, func(i, j int) bool { return list2[i] < list2[j] })

	// Iterates over list 1, adds up total with absolute value from equation.
	total := 0
	for i := 0; i < len(list1); i++ {
		total += int(math.Abs(float64(list1[i] - list2[i])))
	}

	fmt.Println(total)
}
