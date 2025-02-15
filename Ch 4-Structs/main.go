package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	sender := mToSend.sender
	recipient := mToSend.recipient
	if sender.name == "" || sender.number == 0 || recipient.name == "" || recipient.number == 0 {
		return false
	}
	return true
}
