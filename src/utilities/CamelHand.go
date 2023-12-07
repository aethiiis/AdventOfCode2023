package utilities

type Hand struct {
	Force int
	Cards []int
	Bid   int
}

// Compare ComparePart1 Compare Renvoie true si hand est plus grand et false s'il est plus petit
func (hand Hand) Compare(otherHand Hand) bool {
	if hand.Force > otherHand.Force {
		return true
	} else if hand.Force < otherHand.Force {
		return false
	} else if hand.Force == otherHand.Force {
		for i := range hand.Cards {
			if hand.Cards[i] > otherHand.Cards[i] {
				return true
			} else if hand.Cards[i] < otherHand.Cards[i] {
				return false
			} else {
				// On passe à la carte suivante
			}
		}
	}
	panic("Même carte ???")
}
