package main

import (
	"fmt"

	golrn "github.com/lenkite/learn/golrn/internal"
)

type Bingo int
type Tringo int

type Fruit struct {
	name string
}

type Banana Fruit

type BingoExact interface {
	Bingo
}

type BingoApprox interface {
	~int
}

type FruitExactConstraint interface {
	Fruit
}

type GenericBingo[T comparable] struct {
	myfield T
}

func main() {
	fmt.Println("# Learning Go")
	// golrn.DemoArraysAndSlices()
	// golrn.DemoMaps()
	// golrn.DemoStructs()
	//golrn.DemoControlStructures()
	//golrn.DemoErrors()
	golrn.DemoGenerics()
	//golrn.DemoGoRoutines()

}
