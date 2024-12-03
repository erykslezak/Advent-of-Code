package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	total := 0
	// Define the pattern to search for `mul(x,y)`
	pattern := `mul\((\d+),(\d+)\)`
	pat := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matches := pat.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			x, _ := strconv.Atoi(match[1])
			y, _ := strconv.Atoi(match[2])
			total += mul(x, y)
		}
	}
	fmt.Println(total)
}

func mul(x, y int) int {
	return x * y
}
