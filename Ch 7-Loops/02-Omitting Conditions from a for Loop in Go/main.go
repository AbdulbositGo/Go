package main

func maxMessage(thresh int) int {
	total := 0
	for i := 0; ; i++ {
		if total+(100+i) > thresh {
			return i
		}
		total += 100 + i
	}
}
