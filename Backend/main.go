package main

import (
    "fmt"
    "Backend/cards"
)

func main() {
    // Generate a new deck
    deck := cards.GenerateDeck()
    fmt.Println("Original Deck:", deck)

    // Shuffle the deck
    shuffledDeck := cards.ShuffleDeck(deck)
    fmt.Println("\nShuffled Deck:", shuffledDeck)

    // Cut the deck using a player's input
    var playerInput int
    fmt.Println("\nEnter an integer for deck cutting:")
    fmt.Scan(&playerInput)
    cutDeck := cards.CutDeck(shuffledDeck, playerInput)
    fmt.Println("\nDeck after Cutting:", cutDeck)
}
