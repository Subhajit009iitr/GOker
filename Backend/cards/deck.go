package cards

// GenerateDeck initializes and returns a slice containing all 52 cards
func GenerateDeck() []string {
    suits := []string{"D", "H", "S", "C"} // Diamonds, Hearts, Spades, Clubs
    ranks := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

    var deck []string
    for _, suit := range suits {
        for _, rank := range ranks {
            deck = append(deck, suit+rank)
        }
    }
    return deck
}