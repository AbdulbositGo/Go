package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	costs := [3]int{}
	cost := 0

	for i := 0; i < len(messages); i++ {
		cost += len(messages[i])
		costs[i] = cost
	}

	return messages, costs
}
