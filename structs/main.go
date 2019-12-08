package main

import "fmt"

type contactInfo struct {
	email      string
	postalCode string
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Johnson",
		contactInfo: contactInfo{
			email:      "jim.johnson@gmail.com",
			postalCode: "E3S1B5",
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
