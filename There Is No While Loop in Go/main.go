package main

func getMaxMessageToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessageToSend := 0
	balance := float64(maxCostInPennies)

	for balance >= actualCostInPennies {
		balance -= actualCostInPennies
		actualCostInPennies *= costMultiplier
		maxMessageToSend++
	}

	return maxMessageToSend
}
