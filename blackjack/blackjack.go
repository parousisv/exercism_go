package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var card_value int
    switch card{
        case "ace":
    		card_value = 11
        case "two":
    		card_value = 2
        case "three":
    		card_value = 3
        case "four":
    		card_value = 4
        case "five":
    		card_value = 5
        case "six":
    		card_value = 6
        case "seven":
    		card_value = 7
        case "eight":
    		card_value = 8
        case "nine":
    		card_value = 9
        case "ten":
    		card_value = 10
        case "jack":
    		card_value = 10
        case "queen":
    		card_value = 10
        case "king":
    		card_value = 10
        default:
    		card_value = 0
    }
	return card_value
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var returned_value string
    switch (true){
        case card1 == "ace" && card2 == "ace":
    		returned_value = "P"
        case ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) < 10:
    		returned_value = "W"
        case ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) >= 10:
    		returned_value = "S"
        case ParseCard(card1) + ParseCard(card2) >= 17 && ParseCard(card1) + ParseCard(card2) <= 21:
    		returned_value = "S"
        case ParseCard(card1) + ParseCard(card2) >= 12 && ParseCard(card1) + ParseCard(card2) <= 16 && ParseCard(dealerCard) < 7:
    		returned_value = "S"
        case ParseCard(card1) + ParseCard(card2) >= 12 && ParseCard(card1) + ParseCard(card2) <= 16 && ParseCard(dealerCard) >= 7:
    		returned_value = "H"
        case ParseCard(card1) + ParseCard(card2) <= 11:
    		returned_value = "H"
        default:
    		returned_value = "0"
    }
	return returned_value
	panic("Please implement the FirstTurn function")
}
