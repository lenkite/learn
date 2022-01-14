package golrn

import "fmt"

func DemoStructs() {
	fmt.Println("## Structs")
	StructConversions()
}
func StructConversions() {
	fmt.Println("### StructBasics")
	type firstPerson struct {
		name string
		age  int
	}
	type secondPerson struct {
		name string
		age  int
	}
	f1 := firstPerson{
		"bob", 20,
	}
	fmt.Printf("f1: %#v\n", f1)
	var s1 secondPerson
	s1 = secondPerson(f1)
	fmt.Printf("s1: %#v\n", s1)

	fmt.Println("Anonymous structs add a small twist to this: if two struct variables are being compared and at \n" +
		" least one of them has a type that’s an anonymous struct, you can com‐ pare them without a type conversion \n" +
		" if the fields of both structs have the same names, order, and types. You can also assign between named and \n" +
		" anonymous struct types if the fields of both structs have the same names, order, and types")
	var t1 struct {
		name string
		age  int
	}
	t1 = f1
	fmt.Printf("t1: %#v\n", t1)
}

