package main

func getMaxMessageToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessageToSend := 0
	for float64(maxCostInPennies) <= actualCostInPennies {
		maxMessageToSend++
		actualCostInPennies*=costMultiplier
	}
	return maxMessageToSend
}
