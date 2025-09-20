package shuffler

type Suit string

const (
	Clubs		Suit = "Clubs"
	Diamonds	Suit = "Diamonds"
	Hearts		Suit = "Hearts"
	Spades		Suit = "Spades"
	None		Suit = "None"
)

type Rank struct {
	Value	int
	Name	string
}

var (
	Joker	= Rank{1, "Joker"}
	Two		= Rank{2, "Two"}
	Three	= Rank{3, "Three"}
	Four	= Rank{4, "Four"}
	Five	= Rank{5, "Five"}
	Six		= Rank{6, "Six"}
	Seven	= Rank{7, "Seven"}
	Eight	= Rank{8, "Eight"}
	Nine	= Rank{9, "Nine"}
	Ten		= Rank{10, "Ten"}
	Jack	= Rank{11, "Jack"}
	Queen	= Rank{12, "Queen"}
	King	= Rank{13, "King"}
	Ace		= Rank{14, "Ace"}
)

type Card struct {
	suit	Suit `json:"suit"`
	rank	Rank `json:"rank"`
	
}

type Deck struct {
	cards	[]Card	`json:"cards"`
}

type DeckType string

const (
	Standard		DeckType = "Standard"
	StandardJokers	DeckType = "StandardJokers"
	Euchre			DeckType = "Euchre"
	CasinoBlackJack	DeckType = "CasinoBlackJack"

)

func makeStandardDeck() Deck {
	deck := Deck{
		cards: []Card{
			{suit: Clubs,rank: Two,}, {suit: Clubs,rank: Three,}, {suit: Clubs,rank: Four,}, {suit: Clubs,rank: Five,}, {suit: Clubs,rank: Six,}, 
			{suit: Clubs,rank: Seven,}, {suit: Clubs,rank: Eight,}, {suit: Clubs,rank: Nine,}, {suit: Clubs,rank: Ten,}, 
			{suit: Clubs,rank: Jack,}, {suit: Clubs,rank: Queen,}, {suit: Clubs,rank: King,},{suit: Clubs,rank: Ace,},
			
			{suit: Diamonds,rank: Two,}, {suit: Diamonds,rank: Three,}, {suit: Diamonds,rank: Four,}, {suit: Diamonds,rank: Five,}, {suit: Diamonds,rank: Six,}, 
			{suit: Diamonds,rank: Seven,}, {suit: Diamonds,rank: Eight,}, {suit: Diamonds,rank: Nine,}, {suit: Diamonds,rank: Ten,}, 
			{suit: Diamonds,rank: Jack,}, {suit: Diamonds,rank: Queen,}, {suit: Diamonds,rank: King,}, {suit: Diamonds,rank: Ace,}, 
			
			{suit: Hearts,rank: Two,}, {suit: Hearts,rank: Three,}, {suit: Hearts,rank: Four,}, {suit: Hearts,rank: Five,}, {suit: Hearts,rank: Six,}, 
			{suit: Hearts,rank: Seven,}, {suit: Hearts,rank: Eight,}, {suit: Hearts,rank: Nine,}, {suit: Hearts,rank: Ten,}, 
			{suit: Hearts,rank: Jack,}, {suit: Hearts,rank: Queen,}, {suit: Hearts,rank: King,}, {suit: Hearts,rank: Ace,}, 
			
			{suit: Spades,rank: Two,}, {suit: Spades,rank: Three,}, {suit: Spades,rank: Four,}, {suit: Spades,rank: Five,}, {suit: Spades,rank: Six,}, 
			{suit: Spades,rank: Seven,}, {suit: Spades,rank: Eight,}, {suit: Spades,rank: Nine,}, {suit: Spades,rank: Ten,}, 
			{suit: Spades,rank: Jack,}, {suit: Spades,rank: Queen,}, {suit: Spades,rank: King,}, {suit: Spades,rank: Ace,}, 
		},
	}
	return deck

}

func (d *Deck) AddCards(newCards ...Card) {
	d.cards = append(d.cards, newCards...)
}

func (d *Deck) addJokers(num int) {
	cards := make([]Card, num)
	for i := 0; i< num; i++{
		cards[i] = Card{suit: None, rank: Joker}
	}

	d.AddCards(cards...)
}

func (d *Deck) RemoveCards(deleteOne bool, deleteCards ...Card) {
	for i, card := range d.cards {
		for _, deleteCard := range deleteCards {
			if deleteCard == card{
				d.cards = append(d.cards[:i], d.cards[i+1:]...)
				if deleteOne{
					return
				}
			}
		}
	}
}

func (d *Deck) removeCardsByRank(ranks []Rank) {
	for _, rank := range ranks{
		for i, card := range d.cards {
			if card.rank == rank{
				d.cards = append(d.cards[:i], d.cards[i+1:]...)
			}
		}
	}
}


func CreateDeck(deckType DeckType) Deck {
	switch deckType {
		case "Standard":
			return makeStandardDeck()
		case "StandardJokers":
			deck := makeStandardDeck()
			deck.addJokers(2)
			return deck
		case "Euchre":
			deck := makeStandardDeck()
			ranks := []Rank{
				Two, Three, Four, Five, Six, Seven, Eight,
			}
			deck.removeCardsByRank(ranks)

			return deck
		case "CasinoBlackJack":
			var deck Deck
			for i := 0; i < 8; i++ {
				newDeck := makeStandardDeck()
				deck.cards = append(deck.cards, newDeck.cards...)
			}

			return deck
		default:
			return Deck{}
	}
}
