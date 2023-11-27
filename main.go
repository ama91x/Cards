package main

import "fmt"

func main() {
	cards := readAFile("my_Cards")

	cards.PrintDeck()
	cards.shuffle()
	fmt.Println("--------------------")
	cards.PrintDeck()
}
