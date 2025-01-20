package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	consumer, err := sendSMS(msgToCustomer)
	spouse, err1 := sendSMS(msgToSpouse)

	if err != nil {
		return 0, err
	}
	if err1 != nil {
		return 0, err1
	}

	return consumer + spouse, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLength = 25
	const costPerChar = 2
	if len(message) > maxTextLength {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLength)
	}
	return len(message) * costPerChar, nil
}
