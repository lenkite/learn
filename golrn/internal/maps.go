package golrn

import "fmt"

func DemoMaps() {
	fmt.Println("## Maps")
	MapBasics()
	MapCommaOkIdiom()
	MapAsSet()
}

func MapBasics() {
	fmt.Println("### MapBasics")
	var nilMap map[string]int
	fmt.Println("nilMap,len: ", nilMap, len(nilMap))
	emptyMap := map[string]int{}
	fmt.Println("emptyMap,len: ", emptyMap, len(emptyMap))
	teams := map[string][]string{
		"Orcas": {"Fred", "Ralph", "Bijou"}, "Lions": {"Sarah", "Peter", "Billie"}, "Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Println("teams: ", teams)
	fmt.Println("Maps are not comparable. You can check if they are equal to nil\n" +
		", but you cannot check if two maps have identical keys and values using == or differ using !=.")
	newKittens := append(teams["Kittens"], "saucy")
	fmt.Println("newKittens:", newKittens)

	fmt.Println("* Incrementing map keys")

	totalWins := map[string]int{}
	fmt.Println("1. totalWins: ", totalWins)
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++
	fmt.Println(totalWins["Kittens"])
	totalWins["Lions"] = 3
	fmt.Println(totalWins["Lions"])
	fmt.Println("2. totalWins: ", totalWins)
}

func MapCommaOkIdiom() {
	fmt.Println()
	fmt.Println("### MapCommaOkIdiom")
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Println("Rather than assign the result of a map read to a single variable, with the comma ok idiom you \n" +
		"assign the results of a map read to two variables. The first gets the value associated with the key. \n" +
		"The second value returned is a bool.\n" +
		" It is usually named ok. If ok is true, the key is present in the map. If ok is false, the key is not present.")
	fmt.Println("m: ", m)
	v, ok := m["hello"]
	fmt.Println("1. lookup hello: v,ok: ", v, ok)

	v, ok = m["world"]
	fmt.Println("1. lookup world: v,ok: ", v, ok)

	v, ok = m["bingo"]
	fmt.Println("1. lookup bingo: v,ok: ", v, ok)

	fmt.Println("2. deleting world key")
	delete(m, "world")
	v, ok = m["world"]
	fmt.Println("2. lookup world: v,ok: ", v, ok)

}
func MapAsSet() {
	fmt.Println()
	fmt.Println("### MapAsSet")
	fmt.Println(" Use the key of the map for the type that you want to put into the set and use a bool for the value")
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals{
		intSet[v] = true
	}
	fmt.Println("vals, len(vals), len(intSet): ", vals, len(vals), len(intSet))
	fmt.Println("intSet[5]:", intSet[5])
	fmt.Println("intSet[500]:", intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}
}
