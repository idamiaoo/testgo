package hw4

func canCompleteCircuit(gas []int, cost []int) int {
	for i := range gas {
		if gas[i] > cost[i] || (gas[i] == cost[i] && i == len(gas)-1) {
			if canCompleteCircuit1(i, gas, cost) {
				return i
			}
		}

	}
	return -1
}

func canCompleteCircuit1(index int, gas []int, cost []int) bool {
	var leftGas int
	var cnt int

	for cnt < len(gas) {
		leftGas += gas[index]
		if leftGas < cost[index] {
			return false
		}
		leftGas -= cost[index]
		index++
		if index >= len(gas) {
			index = 0
		}
		cnt++
	}
	return true
}
