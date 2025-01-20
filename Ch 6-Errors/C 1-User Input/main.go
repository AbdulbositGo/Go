package main

import "errors"

func validateStatus(status string) error {
	status_len := len(status)
	if status_len < 1 {
		return errors.New("status can not be empty")
	}
	if status_len > 140 {
		return errors.New("status exceeds 140 characters")
	}
	return nil

}
