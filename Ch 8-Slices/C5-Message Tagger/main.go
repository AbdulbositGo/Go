package main

import "strings"

type sms struct {
	id      string
	content string
	tag     []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, message := range messages {
		messages[i].tag = tagger(message)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}
	return tags
}
