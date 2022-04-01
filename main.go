// Struct -> dictionary in python in very high level
// Collection of different properties which are related together.
package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// // Embedded struct
// type person struct {
// 	firstName string
// 	lastName  string
// 	contact   contactInfo
// }

// another way:
type person struct {
	firstName   string
	lastName    string
	contactInfo // Skip the field name.
}

// func main() {
// 	// alex := person{firstName: "Alex", lastName: "Anderson"}
// 	// fmt.Println(alex)

// 	// Default values:
// 	//	string = ""
// 	//	int & float = 0
// 	//	bool = false
// 	var alex person
// 	alex.firstName = "Alex"
// 	alex.lastName = "Anderson"
// 	fmt.Println(alex)
// 	fmt.Printf("%+v", alex) // {firstName: lastName:}

// }

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	// fmt.Printf("%+v", jim)
	// // {firstName:Jim lastName:Party contact:{email:jim@gmail.com zipCode:9400}}

	// jim.updateName("jimmy")
	// jim.print()

	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	// ---------- IMPORTANT -----------
	// Go shortcut for passing by reference.
	jim.updateName("jimmy")
	jim.print() // {firstName:jimmy lastName:Party contactInfo:{email:jim@gmail.com zipCode:9400}}

	// Why it worked by passing person value itself?
	// Due to shortcut: We have defined the receiver as pointer although we are passing value.
	// It really depends on what type we have in receiver.
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

// Print Output:
// jim.updateName("jimmy") --> Pass by value
// jim.print()
// {firstName:Jim lastName:Party contactInfo:{email:jim@gmail.com zipCode:9400}}
// ?? NEED POINTERS ??
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName

// p.print()
// {firstName:jimmy lastName:Party contactInfo:{email:jim@gmail.com zipCode:9400}}
// }

// Go -> Pass by value by default.

// Pass by reference
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

// ------------ IMPORTANT ----------------
// * in front of a type : type description. which means a pointer to a "person"
// * in front of a actual pointer : it is an operator. It means we want to manipulate the value the pointer is referencing to.

// * <pointer> -> turns to value

// turn address to a value: *address
// turn value to address: &value

// ------------------------- STRING SLICE - PASS BY REFERENCE -----------------
// package main

// import "fmt"

// func main() {
// 	mySlice := []string{"Hi", "there", "how", "are", "you"}
// 	updateSlice(mySlice)
// 	fmt.Println(mySlice) // [Bye there how are you] // why???????? Although the slice is passed by value, the ptr still points to the original array.
// }

// func updateSlice(s []string) {
// 	s[0] = "Bye"
// }

// Arrays:
//		- Primitive data structure
//		- Can't be resized
//		- Rarely used directly

// Slice
//		- Can grow & shrink
//		- Used 99% of the time for lists of elements

// When you create a slice, Go internally creates two different data structures.
//		1) Slice
//			- It has three elements inside of it.
//			- 		 i) Pointer - pointer to underlying array that contain actual elements.
//					ii) Capacity - how many elements it can contain at present.
//					iii) Length - how many elements currently exists in the slice.
//		2) Array

// ------------ IMPORTANT ----------------
// Value Types - use pointers to change the values
// 	1) int
//	2) float
//	3) string
//	4) bool
//	5) structs

// ------------ IMPORTANT ----------------
// Reference Types - don't worry about pointers
//	1) slices
//	2) maps
//	3) channels
//	4) pointers
///	5) functions
