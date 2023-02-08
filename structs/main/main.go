package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	human1 := person{firstName: "John",
		lastName: "Doe",
		contact: contactInfo{
			email:   "john.doe@gmail.com",
			zipCode: 95112,
		},
	}
	(&human1).updateName("Joe")
	human1.print()
}

func (p person) print() {
	fmt.Println(p)
}

// will not update
// go is pass by value
func (pPointer *person) updateName(newName string) {
	(*pPointer).firstName = newName
}
