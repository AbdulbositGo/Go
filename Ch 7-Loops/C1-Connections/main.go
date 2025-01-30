package main

import "fmt"

func countConnections(groupSize int) int {
	total := 0
	for i := 1; i < groupSize; i++ {
		total += i
	}
	return total
}

func main() {
	fmt.Println(countConnections(5))
}
