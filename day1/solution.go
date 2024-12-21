package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// var count int

func Solution() {

	var List1 []int
	var List2 []int
	var err1 error
	var err2 error
	var value1 int
	var value2 int
	// count = 1
	filename := "day1/input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening %v file: %v", filename, err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lists := strings.Fields(line)

		value1, err1 = strconv.Atoi(lists[0])
		value2, err2 = strconv.Atoi(lists[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("error converting input value string to int")
		}
		List1 = append(List1, value1)
		List2 = append(List2, value2)
	}

	if len(List1) != len(List2) && len(List1) != 1000 {
		log.Fatalf("Lists do not have the same amount of integers")
	}

	fmt.Println("Solution Part2", countDouble(List1, List2))
	fmt.Println("Solution Part1", recMatchsmallest(List1, List2))
}

func recMatchsmallest(List1 []int, List2 []int) int {

	if len(List1) == 0 || List1 == nil {
		return 0
	}

	var distance int
	var minValue1 int
	var minValue2 int
	var minIndex1 int
	var minIndex2 int

	minValue1 = List1[0]
	for i := range List1 {
		if List1[i] < minValue1 {
			minValue1 = List1[i]
			minIndex1 = i

		}
	}
	minValue2 = List2[0]
	for j := range List2 {
		if List2[j] < minValue2 {
			minValue2 = List2[j]
			minIndex2 = j
		}
	}
	// fmt.Printf("%d,     List 1: %d      List2: %d      Distance:%d   \n", count, minValue1, minValue2, getDifference(minValue1, minValue2))
	// count++
	List1 = append(List1[:minIndex1], List1[minIndex1+1:]...)
	List2 = append(List2[:minIndex2], List2[minIndex2+1:]...)

	distance = getDifference(minValue1, minValue2) + recMatchsmallest(List1, List2)

	return distance
}

func getDifference(value1 int, value2 int) int {
	if value1 <= value2 {
		return value2 - value1
	} else {
		return value1 - value2
	}
}

func countDouble(list1 []int, list2 []int) int {
	var doubles int

	for i := range list1 {
		numberCount := 0
		for j := range list2 {
			if list1[i] == list2[j] {
				numberCount++
			}
		}
		doubles += list1[i] * numberCount
	}
	return doubles
}
