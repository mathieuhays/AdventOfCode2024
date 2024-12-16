package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/mathieuhays/AdventOfCode2024/internals/utils"
	"io"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Advent Of Code, Day 01")

	if err := utils.Validate(); err != nil {
		log.Fatal(err)
	}

	part, _ := utils.GetPart()

	reader, err := utils.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	if part == 1 {
		err, diff := getTotalDiff(reader)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Total diff: %v\n", diff)
	} else if part == 2 {
		err, score := getSimilarityScore(reader)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Similarity Score: %v\n", score)
	}

	fmt.Printf("Executed in %v\n", time.Now().Sub(start))
}

func getSimilarityScore(reader io.Reader) (error, int) {
	err, left, right := inputToLists(reader)

	if err != nil {
		return fmt.Errorf("inputToLists: %s", err), 0
	}

	err, rightCounts := getListCounts(right)
	if err != nil {
		return fmt.Errorf("getListCounts: %s", err), 0
	}

	var scores []int

	for _, v := range left {
		if count, ok := rightCounts[v]; ok {
			scores = append(scores, count*v)
		}
	}

	return nil, reduce(scores)
}

func getListCounts(list []int) (error, map[int]int) {
	counts := map[int]int{}

	for _, v := range list {
		if _, ok := counts[v]; ok {
			counts[v]++
		} else {
			counts[v] = 1
		}
	}

	return nil, counts
}

func getTotalDiff(reader io.Reader) (error, int) {
	err, left, right := inputToLists(reader)

	if err != nil {
		return fmt.Errorf("inputToLists: %s", err), 0
	}

	slices.Sort(left)
	slices.Sort(right)

	err, diffs := listsToDiff(left, right)
	if err != nil {
		return fmt.Errorf("listsToDiff: %s", err), 0
	}

	return nil, reduce(diffs)
}

func reduce(list []int) int {
	total := 0

	for _, v := range list {
		total += v
	}

	return total
}

func listsToDiff(a []int, b []int) (error, []int) {
	if len(a) != len(b) {
		return errors.New("invalid inputs (diff length)"), []int{}
	}

	diff := make([]int, len(a))

	for i := 0; i < len(a); i++ {
		localDiff := math.Abs(float64(a[i]) - float64(b[i]))
		diff = append(diff, int(localDiff))
	}

	return nil, diff
}

func inputToLists(reader io.Reader) (error, []int, []int) {
	scanner := bufio.NewScanner(reader)
	left := []int{}
	right := []int{}

	for scanner.Scan() {
		err, l, r := lineToInts(scanner.Text())
		if err == nil {
			left = append(left, l)
			right = append(right, r)
		}
	}

	return nil, left, right
}

func lineToInts(input string) (error, int, int) {
	parts := strings.Split(input, " ")
	numbers := []int{}

	for _, p := range parts {
		if p == " " {
			continue
		}

		i, err := strconv.Atoi(p)
		if err != nil {
			continue
		}

		numbers = append(numbers, i)
	}

	if len(numbers) != 2 {
		return errors.New("invalid"), 0, 0
	}

	return nil, numbers[0], numbers[1]
}
