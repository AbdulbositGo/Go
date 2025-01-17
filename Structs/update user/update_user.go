package main

import "fmt"

type User struct {
	name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name, membershipType string) User {
	user := User{
		name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: 100,
		},
	}
	if membershipType == "premium" {
		user.Membership.MessageCharLimit = 1000
	}
	return user
}

func main() {
	// Example usage of the newUser function
	user1 := newUser("Alice", "premium")
	user2 := newUser("Bob", "basic")

	// Print the details of each user
	fmt.Printf("User: %s, Membership: %s, Limit: %d\n", user1.name, user1.Type, user1.MessageCharLimit)
	fmt.Printf("User: %s, Membership: %s, Limit: %d\n", user2.name, user2.Type, user2.MessageCharLimit)
	fmt.Println(user1)
}
