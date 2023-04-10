package main

import (
	"fmt"
)

/*
# Struct is a Custom data type.
# Struct is like "Class" or "BluePrint" Because anyone can use that then make their
  own custom data by using those fields.
*/

// Here we make a struct [Custom dataType]
type StructPointer struct {
	val string
	id  int
}

// we Embedd the "StructPointer" with StrPointer,StructPointer is a recevier.
func (s StructPointer) StrPointer() {
	fmt.Println("StructPointer", s.val, s.id)
}

/*
#Call By Value.
#When a function or method is called with a value argument, the function receives a
copy of the value, and any changes made to that value inside the function or method
are only made to the copy, not the original value
*/
func (s StructPointer) UpdateId(id int) {
	s.id = id
	fmt.Println("s", s)
}
/*
#Call By Reference
# we getting a original value rather than a copy of the value.
# when we putting a receiver as a pointer receiver.
*/
func (s *StructPointer)UpdateIdWithPtrReceiver(id int)  {
	/*
	# Here we declare a "struct" as pointer but the the fields all are normal
	  so declare a value directly to those fileds.
	*/
	s.id = id
	fmt.Println("s", s)
}

func main() {
	// Here we intiate the "StructPointer" as well as giving value for that in str varible.
	str := StructPointer{val: "one", id: 1}
	str.UpdateId(5)

	str.StrPointer()
	// After that we again make a "StructPointer" as well as giving value for that in str2 varible.
	str2 := StructPointer{val: "two", id: 2}
	str2.UpdateIdWithPtrReceiver(10)
	str2.StrPointer()
	// Both aren't connected.
}
