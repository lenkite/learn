package golrn

import "fmt"

func DemoGoRoutines() {
	fmt.Println(" # DemoGoRoutines")
	SquareJob()
	SelectDemo()
}

func SelectDemo() {
}

func doSelect(ch1 <-chan int, ch2 <-chan int, ch3 chan<- int, ch4 <-chan int) {
	// The select stmt allows a go-routine to read from or write to a set of multiple channels
	// looks a great deal like a blank switch.
	// EAch case in a select is a read or write to a channel.
	// if a read or write is possible for a case, it is executed with the body of the case.
	x := 10
	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	case ch3 <- x:
		fmt.Println("Wrote:", x)
	case <-ch4:
		fmt.Println("Got val fomm ch4, ignored.")
	}
}

func SquareJob() {
	vals := makeRange(1, 999)
	a := make(chan int)
	b := make(chan int)
	go func() {
		for r := range b {
			fmt.Printf("(SquareJob) Got %v from: %v\n", r, b)
		}
	}()
	// a := <-ch, ch <- b
	readProcessWrite(a, func(i int) int { return i * i }, b)
	for v := range vals {
		fmt.Printf("(SquareJob) Put %v into: %v\n", v, a)
		a <- v
	}
	//readProcessWrite()
}

func readProcessWrite(in <-chan int, process func(int) int, out chan<- int) {
	go func() {
		for v := range in {
			fmt.Printf("(readProcessWrite) Read %v from %v\n", v, in)
			r := process(v)
			out <- r
			fmt.Printf("(readProcessWrite) Wrote %v to %v\n", r, out)
		}
	}()
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
