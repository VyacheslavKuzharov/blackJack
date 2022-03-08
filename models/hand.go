package models

type Hand struct {
	Cards []Card
}

func (h *Hand) Create(hand []Card) []Card {
	var userHand []Card

	userHand = append(userHand, hand...)
	return userHand
}
