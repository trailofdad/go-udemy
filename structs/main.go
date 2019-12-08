package main

import "fmt"

type contactInfo struct {
	email      string
	postalCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Johnson",
		contact: contactInfo{
			email:      "jim.johnson@gmail.com",
			postalCode: "E3S1B5",
		},
	}

	fmt.Printf("%+v", jim)
}
