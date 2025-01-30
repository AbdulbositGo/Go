package main

import "unicode"

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}
	isDigit := false
	isUpper := false
	for _, word := range password {
		if unicode.IsDigit(word) {
			isDigit = true
		} else if unicode.IsUpper(word) {
			isUpper = true
		}
	}
	if isDigit && isUpper {
		return true
	}
	return false
}
