package deck

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Spade})
	fmt.Println(Card{Rank: Four, Suit: Heart})
	fmt.Println(Card{Rank: King, Suit: Diamond})
	fmt.Println(Card{Rank: Ten, Suit: Club})
	fmt.Println(Card{Suit: Joker})

	// Output: 
	// Ace of Spades
	// Four of Hearts
	// King of Diamonds
	// Ten of Clubs
	// Joker

}

func TestNew (t *testing.T) {
	cards := NewCard()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong amount of cards created")
	}
}

func TestDefaultSort (t *testing.T) {
	cards := NewCard(DefaultSort)
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades as first result, instead got:", cards[0])
	}
}

func TestSort (t *testing.T) {
	cards := NewCard(Sort(Less))
	expected := Card{Rank: Ace, Suit: Spade}
	if cards[0] != expected {
		t.Error("Expected Ace of Spades as first result, instead got:", cards[0])
	}
}