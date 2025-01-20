package main

import "fmt"

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type      string
	CharLimit int
}

func (u User) SendMessage(message string, length int) (string, bool) {
	if u.Membership.CharLimit < length {
		return "", false
	}
	return message, true
}

func main() {
	// Example usage
	user := User{
		Name: "Alice",
		Membership: Membership{
			Type:      "premium",
			CharLimit: 1000,
		},
	}

	message := "Hello, this is a test message."
	length := len(message)

	result, canSend := user.SendMessage(message, length)

	if canSend {
		fmt.Printf("Message sent: %s\n", result)
	} else {
		fmt.Println("Message cannot be sent: exceeds character limit.")
	}
}
