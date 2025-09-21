package shuffler

import (
	"math/rand/v2"
	"time"
)

func (d *Deck) BridgeShuffler() {
	cycles := (rand.IntN(3) + 1) * 3
	for _ = range(cycles){
		skip := rand.IntN(2) + 1
		start := rand.IntN(1)
		for i := start; i < len(d.cards); i += skip{
			d.cards = append(d.cards, d.cards[i])
			d.cards = append(d.cards[:i], d.cards[i+1:]...)
		}
	}
}

func (d *Deck) StackShuffler() {
	cycles := (rand.IntN(3) + 1) * 3
	for _ = range(cycles){
		start := rand.IntN(len(d.cards) - 8)
		amount := rand.IntN(7)

		stackCards := make([]Card, amount)
		copy(stackCards, d.cards[start:start+amount])

		d.cards = append(d.cards[:start], d.cards[start+amount:]...)
		d.cards = append(d.cards, stackCards...)
	}
}

func (d *Deck) ConstantStackShuffler(freq int) {
	for {
		d.StackShuffler()
		time.Sleep(time.Duration(freq) * time.Millisecond)
	}
}

func (d *Deck) ConstantBridgeShuffler(freq int) {
	for {
		d.BridgeShuffler()
		time.Sleep(time.Duration(freq) * time.Millisecond)
	}
}