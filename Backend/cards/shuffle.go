package cards

import (
    "math/rand"
    "time"
)

// ShuffleDeck shuffles the cards in the provided deck
func ShuffleDeck(deck []string) []string {
    // Create a new random generator with a time-based source
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    rng.Shuffle(len(deck), func(i, j int) {
        deck[i], deck[j] = deck[j], deck[i]
    })
    return deck
}

// CutDeck simulates a deck cut based on the player's input
func CutDeck(deck []string, n int) []string {
    n = n % 52 // Ensure n is within the bounds of the deck
    if n == 0 {
        return deck // No cutting needed if n is 0
    }
    return append(deck[n:], deck[:n]...) // Move first n cards to the end
}
