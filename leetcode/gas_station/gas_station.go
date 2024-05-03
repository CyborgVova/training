package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	total, index, road := 0, 0, 0

	for i := 0; i < len(gas); i++ {
		share := gas[i] - cost[i]
		road += share
		total += share
		if road < 0 {
			index = i + 1
			road = 0
		}
	}

	if total < 0 {
		return -1
	}

	return index
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}

	fmt.Println(canCompleteCircuit(gas, cost))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
