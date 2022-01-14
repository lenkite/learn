package golrn

import "fmt"

func DemoControlStructures() {
	fmt.Println("## Control Structures")
	Switch()
	BlankSwitch();
}
func Switch() {
	fmt.Println("### Switch")

	words := []string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

}
func BlankSwitch() {
	fmt.Println("### BlankSwitch")
	fmt.Println("you can write a switch state‐ ment that doesn’t specify the value that you’re comparing against. \n" +
		"A blank switch allows you to use any boolean comparison for each case. ")

}
