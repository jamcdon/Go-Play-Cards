package main

import (
	"fmt"
	"go-play-cards/shuffler"
	"time"
)

func main() {
	deck := shuffler.CreateDeck(shuffler.Euchre)
	fmt.Println(deck)
	//deck.StackShuffler()
	//fmt.Println(deck)
	go deck.ConstantBridgeShuffler(500)
	time.Sleep(1 * time.Second)
	fmt.Println(deck)
	time.Sleep(1 * time.Second)
	fmt.Println(deck)
}