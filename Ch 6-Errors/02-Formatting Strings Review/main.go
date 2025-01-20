package main

import (
	"fmt"
)

func getSMSErrorString(cost float64, recipient string) string {
	return fmt.Sprintf("SMS that cost $%.2f to be sent to '%s' can not be sent", cost, recipient)
}
