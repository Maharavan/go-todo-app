package main

import (
	"fmt"
	"strconv"
)

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
			phoneno: 9000000000,
		},
	}
	slices := []string{"Hello", "Good", "Morning"}
	updateSlice(slices)
	fmt.Println(slices)
	// fmt.Printf("%+v", data)
	dataPointer := &data // point to address memory of data
	dataPointer.updateFirstName("Jim")
	data.printcontent() // so print the original value

	// data.updateFirstName("Jim")
	// data.printcontent() without passing passing address ,Go will directly take address of it
}

func (pointer *person) updateFirstName(newString string) { //(*person -> (pointer to person) The *person means that the receiver pointer is a variable that holds the address of a person struct.
	(*pointer).firstname = newString //another way of referencing so called manual reference
	pointer.lastname = "hel"         //auto dereferencing feature from Go
	pointer.contactinfo.emailid = "leooo@gmail.com"
	fmt.Println(len(strconv.Itoa(pointer.contactinfo.phoneno)))
}
func (p person) printcontent() {
	fmt.Printf("%+v", p)
}

// slice have pointer,length and capacity in general so affecting will directly change in general
// Go is pass-by-value language so it copies value to other address since it has pointer so eventhough stored in different address it maps to location where value stored (reference type)
// int ,float,string,bool,structs -> value type , slices,maos,channels,pointers,functions -> reference type
func updateSlice(slice []string) {
	slice[0] = "Amigos"
	slice = append(slice, "Leo")
}
