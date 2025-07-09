package main

import "fmt"

type contactinfo struct {
	emailid string
	phoneno int
}
type person struct {
	firstname string
	lastname  string
	contactinfo
}

func main() {
	var data person
	data = person{
		firstname: "mahar",
		lastname:  "s",
		contactinfo: contactinfo{
			emailid: "mahars@gmail.com",
			phoneno: 9385954437,
		},
	}

	// fmt.Printf("%+v", data)
	data.updateFirstName("Jim")
	data.printcontent()
}

func (p person) updateFirstName(newString string) {
	p.firstname = newString
	fmt.Printf("%+v", p)
}
func (p person) printcontent() {
	fmt.Printf("%+v", p)
}
