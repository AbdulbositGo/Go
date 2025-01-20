package main

func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", 0.0
	}
}

type expense interface {
	cost() float64
}

type email struct {
	body         string
	toAddress    string
	isSubscribed bool
}

type sms struct {
	body          string
	toPhoneNumber string
	isSubscribed  bool
}

type invalid struct{}

func (e email) cost() float64 {
	if e.isSubscribed {
		return float64(len(e.body)) * 0.01
	}
	return float64(len(e.body)) * 0.05
}

func (s sms) cost() float64 {
	if s.isSubscribed {
		return float64(len(s.body)) * 0.05
	}
	return float64(len(s.body)) * 0.1
}
