package golrn

import "fmt"

func DemoArraysAndSlices() {
	fmt.Println(" ## Arrays And Slices")
	ArrayBasics()
	SliceBasics()
	SliceAppends()
	SliceLenCap()
	SliceViaMake()
	SliceExpressions()
	SliceOverlapping()
	SliceCopying()
}

func ArrayBasics() {
	fmt.Println("### ArrayBasics")
	var x1 [3]int
	fmt.Println("x1 = ", x1) // all positions zero initialized.
	var x2 = [...]int{10, 20, 30}
	fmt.Println("x2 = ", x2, ", len(x2) = ", len(x2))
	var y2 = [3]int{10, 20, 30}
	fmt.Println("x2 =  y2 ?", x2 == y2) // == and != can be used to compare arrays.
}

func SliceBasics() {
	fmt.Println()
	fmt.Println("### SliceBasics")
	var x1 = []int{10, 20, 30}
	fmt.Println("x1 = ", x1)
	var x2 = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println("x2 = ", x2)
	var x3 = []int{1, 5: 4, 23, 3: 2, 2: 0, 1: 99, 8: 32, 4}
	fmt.Println("x3 = ", x3)
	fmt.Println("\n" +
		"A slice is the first type we’ve seen that isn’t comparable. \n" +
		"It is a compile-time error to use == to " +
		"see if two slices are identical or != to see if they are different. " +
		"The only thing you can compare a slice with is nil:")
}

func SliceAppends() {
	fmt.Println()
	fmt.Println("### SliceAppends")
	fmt.Println("The built-in append function is used to grow slices. \n" +
		"It is a compile-time error if you forget to assign the value returned from append.")
	var x1 = []int{1, 2, 3}
	x1 = append(x1, 4)
	fmt.Println("x1 = ", x1)
	x1 = append(x1, 5, 6, 7, 8)
	fmt.Println(x1)
	fmt.Println("x1 after multi-append = ", x1)
	fmt.Println("One slice is appended onto another by using the ... operator to expand the source slice into individual values ")
	y1 := []int{20, 30, 40}
	fmt.Println("y1 = ", y1)
	x2 := append(x1, y1...)
	fmt.Println("x2 (after appending y1) = ", x2)
}

func SliceLenCap() {
	fmt.Println()
	fmt.Println("### SliceLenCap")
	fmt.Println("The rules as of Go 1.14 are to double the size of the slice when the capacity is less than 1,024 \n " +
		"and then grow by at least 25% afterward.")
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 60)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 70)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 80)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 90)
	fmt.Println(x, len(x), cap(x))
}
func SliceViaMake() {
	fmt.Println()
	fmt.Println("### SliceViaMake")
	x1 := make([]int, 5)
	fmt.Printf("%-20s: ", "x1")
	fmt.Println(x1, len(x1), cap(x1))
	x1= append(x1, 10)
	fmt.Printf("%-20s: ", "x1 after append 10")
	fmt.Println(x1, len(x1), cap(x1))

	fmt.Println()
	fmt.Println("* Nil vs Zero length size literal")
	var data []int
	fmt.Print("data: ")
	fmt.Println(data, len(data), cap(data))
	fmt.Println("data == nil ?", data == nil)
	var x2 = []int{}
	fmt.Print("x2: ")
	fmt.Println(x2, len(x2), cap(x2))
	fmt.Println("x2 == nil ?", x2 == nil)

}

func SliceExpressions() {
	fmt.Println()
	fmt.Println("## SliceExpressions")
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println("When you take a slice from a slice, you are not making a copy of the data. \n" +
		"Instead, you now have two variables that are sharing memory. \n" +
		"This means that changes to an element in a slice affect all slices that share that element. ")

	f := func() {
		x := []int{1, 2, 3, 4}
		y := x[:2]
		z := x[1:]
		x[1] = 20
		y[0] = 10
		z[1] = 30
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	}
	f()

}
func SliceOverlapping() {
	fmt.Println()
	fmt.Println("## SliceOverlapping")
	x := make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	fmt.Println("initial x,lx,cx:", x, len(x), cap(x))
	y := x[:2]
	fmt.Println("initial y,ly,cy:", y, len(y), cap(y))
	z := x[2:]
	fmt.Println("initial z,lz,sz:", z, len(z), cap(z))
	y = append(y, 30, 40, 50)
	fmt.Println("intermediate y,ly,cy:", y, len(y), cap(y))
	x = append(x, 60)
	fmt.Println("intermediate  x,lx,cx:", x, len(x), cap(x))
	z = append(z, 70)
	fmt.Println("intermediate z,lz,sz:", z, len(z), cap(z))
	fmt.Println("final x,lx,cx:", x, len(x), cap(x))
	fmt.Println("final y,ly,cy:", y, len(y), cap(y))
	fmt.Println("final z,lz,sz:", z, len(z), cap(z))
}
func SliceCopying() {
	fmt.Println()
	fmt.Println("### SliceCopying")
	x := []int{1, 2, 3, 4}
	var y []int
	z := make([]int, 4)
	fmt.Println("1. x,lx,cx,isNil:", x, len(x), cap(x), x == nil)
	fmt.Println("1. y,ly,cy,isNil:", y, len(y), cap(y), y == nil)
	//fmt.Println(" y copied from x: ", y)
	cy := copy(y, x)
	fmt.Println("2. (y,len,cap,isNil,count): ", y, len(y), cap(y), y == nil, cy)
	cz := copy(z, x)
	fmt.Println("2. (z,len,cap,isNil,count): ", z, len(z), cap(z), z == nil, cz)
	//fmt.Println(" y copied from x: ", y)
	x[2] = 30
	x[3] = 40
	fmt.Println("3. (x,len,cap,isNil): ", x, len(x), cap(x), x == nil)
	cz = copy(z, x[2:])
	fmt.Println("3. (z,len,cap,isNil,count): ", z, len(z), cap(z), z == nil, cz)

	x1 := []int{1, 2, 3, 4}
	fmt.Println("4. (x1,len,cap,isNil): ", x1, len(x1), cap(x1), x1 == nil)
	fmt.Println("copy(x1[:3], x1[1:]")
	n1 := copy(x1[:3], x1[1:])
	fmt.Println("5. (x1,len,cap,isNil,count): ", x1, len(x1), cap(x1), x1 == nil, n1)
}
