package golrn

import (
	"fmt"
	"golang.org/x/xerrors"
	"runtime"
)
var ModulePath  = "github.com/lenkite/learn/golrn"

func DemoErrors() {
	fmt.Println("## Errors")
	//CustomErrors()
	//ErrorWrapChain()
	//ErrorWithStackTrace()
	//XError()
}


func xA() error {
	return xerrors.Errorf("bad value: %d", 42)
}
func xB() error {
	return xerrors.Errorf("in xB: %w", xA())
}
func xC() error {
	return xerrors.Errorf("in xC: %w", xB())
}



func CustomErrors() {
	fmt.Println("### CustomErrors")
	ra, erra := sillyFuncA(2)
	if erra != nil {
		fmt.Println("(CustomErrors) sillyFuncA gave err: ", erra)
	} else {
		fmt.Println("(CustomErrors) sillyFuncA gav result: ", ra)
	}

	rb, errb := sillyFuncB(2)
	if errb != nil {
		fmt.Println("(CustomErrors) sillFuncB gave err: ", errb)
	} else {
		fmt.Println("(CustomErrors) sillyFuncB gave result: ", rb)
	}
}

func ErrorWrapChain() {
	fmt.Println("### CustomErrors")
	err := singo()
	fmt.Println("(ErrorWrapChain): " , err)
}

func ErrorWithStackTrace() {
	fmt.Println("### ErrorWithStackTrace")
	fmt.Println("(ErrorWithStackTrace)", stC())
	//fmt.Println(stA())
}

func stA() error {
	return Errorf("problem in A: %d", 42)
}
func stB() error {
	return Errorf("problem in B: %w", stA())
}

func stC() error {
	return Errorf("problem in C: %w", stB())
}

func Errorf(format string, args ...interface{}) error {
	pc := make([]uintptr, 1)
	// Skip 2 levels to get the caller
	n := runtime.Callers(2, pc)
	var _, function string
	var line int
	if n > 0 {
		frames := runtime.CallersFrames(pc[:n])
		frame, _ := frames.Next()
		_, function, line = frame.File, frame.Function, frame.Line
		function = function[len(ModulePath):]
	}
	args = append([]interface{}{function, line}, args...)
	format = "(%s|%d) " + format
	return fmt.Errorf(format, args...)
}
func singo() error {
	err := tringo()
	return fmt.Errorf("in singo: %w", err)
}

func tringo() error {
	err := bingo()
	return fmt.Errorf("in tringo: %w", err)
}

func bingo() error {
	return fmt.Errorf("in bingo: %s", "write failed")
}

type StatusErr struct {
	Status  Status
	Message string
}

func (se StatusErr) Error() string {
	return se.Message
}

func sillyFuncA(a int) (int, error) {
	fmt.Println("--Inside sillyFuncA, a: ", a)
	if a%2 == 0 {
		return 0, StatusErr{
			Status:  NotOdd,
			Message: fmt.Sprintf("Not odd %d", a),
		}
	}
	return a*a, nil
}

func sillyFuncB(b int) (int , *StatusErr) {
	fmt.Println("--Inside sillyFuncB, b: ", b)
	if b%2 == 1 {
		return 0, &StatusErr{
			Status:  NotEven,
			Message: fmt.Sprintf("Not even %d", b),
		}
	}
	return b*b, nil
}

type Status int

const (
	NotOdd Status = iota + 1
	NotEven
)

