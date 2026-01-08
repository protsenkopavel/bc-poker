package main

import (
	"bcpocker/deck"
	"fmt"
)

func main() {
	d := deck.Shuffle(deck.New())
	fmt.Println(d)
}
