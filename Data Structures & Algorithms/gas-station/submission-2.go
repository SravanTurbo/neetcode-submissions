/*
start: 8:22am
idea: greedy
t: O(n) s:O(1)
*/
func canCompleteCircuit(gas []int, cost []int) int {
	totalGas, totalCost  := 0, 0
	for i := 0; i < len(gas); i++ {
		totalGas  += gas[i]
		totalCost += cost[i]
	}

	if totalCost > totalGas { //impossible
		return -1 
	}

	start, fuel := 0, 0
	for i := 0; i < len(gas); i++ {
		fuel += gas[i] - cost[i]
		if fuel < 0 {
			start = i+1
			fuel  = 0
		}
	}

	return start
}
