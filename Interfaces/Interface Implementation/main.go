package main

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

type fullTime struct {
	name   string
	salary int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

func (ft fullTime) getName() string {
	return ft.name
}

func (ft fullTime) getSalary() int {
	return ft.salary
}
