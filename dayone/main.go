package main

import (
	"bufio"
	"cmp"
	"fmt"
	"slices"
	"strconv"

	"aoc/dayone/elf"
	"log"
	"os"
)

// Calorie counting
// https://adventofcode.com/2022/day/1
func main() {
	fmt.Println("Opening input file")
	f, err := os.Open("./dayone/input.txt")

	check("file open", err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	id := 1
	elves := []elf.Elf{}

	currentElf := elf.Elf{Id: id, Calories: 0}

	fmt.Println("Scanning elves")

	for scanner.Scan() {
		txt := scanner.Text()

		if len(txt) > 0 {
			calories, err := strconv.Atoi(txt)
			check("cal convert", err)

			currentElf.AddCall(calories)
			continue
		}

		// fmt.Printf("Elf %d calories: %d\n", currentElf.Id, currentElf.Calories)

		id++
		elves = append(elves, currentElf)
		currentElf = elf.Elf{Id: id, Calories: 0}
	}

	fmt.Println("Elf count", len(elves))
	fmt.Println("Sorting elves")

	slices.SortFunc(elves, compareElves)

	fmt.Println("Highest", elves[0].Id, elves[0].Calories)
	fmt.Println("Highest", elves[1].Id, elves[1].Calories)
	fmt.Println("Highest", elves[2].Id, elves[2].Calories)

	total := 0

	for i := 0; i < 3; i++ {
		total += elves[i].Calories
	}

	fmt.Println("[ANSWER]: Total calories", total)
}

func compareElves(a, b elf.Elf) int {
	return cmp.Compare(b.Calories, a.Calories)
}

func check(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
