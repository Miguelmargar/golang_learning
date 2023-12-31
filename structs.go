package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func runStruct() {

	peter := person{"Peter", "Anderson", contactInfo{"test_one@test.com", 12345}}

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "test_two@test.com",
			zipCode: 12345,
		},
	}

	var john person

	var john_contact contactInfo
	john_contact.email = "test_three@test.com"
	john_contact.zipCode = 12345

	john.firstName = "John"
	john.lastName = "Anderson"
	john.contactInfo = john_contact

	fmt.Printf("%+v"+"\n", alex)
	fmt.Println(peter)
	fmt.Println(john)

	peterPointer := &peter // memory pointer
	peterPointer.updateName("Pete")
	peter.print()

	// sintactic sugar for the above in peterPointer no need to create the memory pointer
	john.updateName("Johnny")
	john.print()
}

// *person is a type of pointer to a person - structs are pass by values and slices are by reference.
func (pointerToPerson *person) updateName(newFirstName string) {

	// pointer to value of the memory location for pointerToPerson
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {

	fmt.Printf("%+v"+"\n", p)
}
