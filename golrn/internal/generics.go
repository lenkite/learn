package golrn

import "fmt"

func DemoGenerics() {
	fmt.Println("## Generics")
	//CustomErrors()
	//ErrorWrapChain()
	//ErrorWithStackTrace()
	//DemoSumIntsOrFloats()
	//DemoStringify()
	DemoConcatTo()

}

// ListHead is the head of a linked list.
type ListHead[T any, U any] struct {
	head *ListElement[T, U]
}

// ListElement is an element in a linked list with a head.
// Each element points back to the head.
type ListElement[T any, U any] struct {
	next *ListElement[U, T]
	vala T
	valb U
	// Using ListHead[T] here is OK.
	// ListHead[T] refers to ListElement[T] refers to ListHead[T].
	// Using ListHead[int] would not be OK, as ListHead[T]
	// would have an indirect reference to ListHead[int].
	head *ListHead[U, T]
}

func DemoConcatTo() {
	v := []Animal{{"Dog", 10}, {"Cat", 9}}
	s := ConcatTo(v, v)
	fmt.Printf("DemoConcatTo: s: %v\n", s)
}

var _ fmt.Stringer = Animal{}       // Verify that T implements I.
var _ fmt.Stringer = (*Animal)(nil) // Verify that *T implements I.
var _ Plusser = Animal{}            // Verify that T implements I.
var _ Plusser = (*Animal)(nil)      // Verify that *T implements I.

// Plusser is a type constraint that requires a Plus method.
// The Plus method is expected to add the argument to an internal
// string and return the result.
type Plusser interface {
	Plus(string) string
}

// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func (a Animal) Plus(s string) string {
	return a.Name + ":" + s
}

func DemoStringify() {
	fmt.Println("DemoStringify")
	v := []Animal{{"Dog", 10}, {"Cat", 9}}
	fmt.Println("DemoStringify", Stringify(v))
}

// Stringify calls the String method on each element of s,
// and returns the results.
func Stringify[T fmt.Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// ConcatTo takes a slice of elements with a String method and a slice
// of elements with a Plus method. The slices should have the same
// number of elements. This will convert each element of s to a string,
// pass it to the Plus method of the corresponding element of p,
// and return a slice of the resulting strings.
func ConcatTo[S fmt.Stringer, P Plusser](s []S, p []P) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = p[i].Plus(v.String())
	}
	return r
}
func DemoSumIntsOrFloats() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Printf("Generic Sums: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))
}

type Number interface {
	int64 | float64
}

// SumNumbers sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
