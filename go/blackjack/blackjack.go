package blackjack


// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {

	switch {

		case card == "ace": {
			return  11
		}
		case card == "two":{
			return  2
		}
		case card == "three":{
			return  3
		}
		case card == "four":{
			return  4
		}
		case card == "five":{
			return  5
		}
		case card == "six":{
			return  6
		}
		case card == "seven":{
			return  7
		}
		case card == "eight":{
			return  8
		}
		case card == "nine":{
			return  9
		}
		case card == "ten" ||card == "jack" || card == "queen" || card == "king":{
			return  10
		}
	default:{
		return 0
	}
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	dealerValue := ParseCard(dealerCard)
	sum := value1 + value2

switch {

	// Pair of aces → always split
case value1 == 11 && value2 == 11:
	return "P"

	// Blackjack
case sum == 21 && dealerValue != 10 && dealerValue != 11:
	return "W"

case sum == 21:
	return "S"

	// Sum 17–20 → stand
case sum >= 17 && sum <= 20:
	return "S"

	// Sum 12–16 → hit if dealer >= 7
case sum >= 12 && sum <= 16 && dealerValue >= 7:
	return "H"

	// Sum 12–16 → otherwise stand
case sum >= 12 && sum <= 16:
	return "S"

	// Sum <= 11 → hit
case sum <= 11:
	return "H"

default:
	return "ERRO"
}

}
