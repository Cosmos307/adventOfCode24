package main

import (
	"fmt"

	"github.com/Cosmos307/adventOfCode24/day1"
)

func main() {
	var input int

	fmt.Print("Type a number: ")
	fmt.Scan(&input)

	switch input {
	case 1:
		day1.Solution()

	}
}
