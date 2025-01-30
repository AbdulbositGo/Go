package main

func indexOfFirstBadWord(msg, badWords []string) int {

	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}
